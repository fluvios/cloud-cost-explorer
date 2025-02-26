// Golang Service - azure_service.go (Azure Cost Fetching)
package azure

import (
	"context"
	"time"
	"log"
	"os"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // No password set
	DB:       0,  // Default DB
})

// FetchAzureCloudCosts fetches cost data from Azure Cost Management API
func FetchAzureCloudCosts(startDate, endDate string) (map[string]interface{}, error) {
	ctx := context.TODO()
	cacheKey := "azure_costs:" + startDate + "_" + endDate

	// Check Redis cache first
	cachedData, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Returning cached Azure cost data")
		return map[string]interface{}{"cached": true, "data": cachedData}, nil
	} else if !errors.Is(err, redis.Nil) {
		log.Printf("Error retrieving from cache: %v", err)
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Printf("Error creating Azure credential: %v", err)
		return nil, err
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return nil, errors.New("AZURE_SUBSCRIPTION_ID environment variable is not set")
	}

	client, err := armcostmanagement.NewQueryClient(cred, nil)
	if err != nil {
		log.Printf("Error creating Azure cost management client: %v", err)
		return nil, err
	}

	req := armcostmanagement.QueryDefinition{
		Type: to.Ptr(armcostmanagement.ExportTypeActualCost),
		Timeframe: to.Ptr(armcostmanagement.TimeframeCustom),
		TimePeriod: &armcostmanagement.QueryTimePeriod{
			From: to.Ptr(startDate),
			To:   to.Ptr(endDate),
		},
	}

	resp, err := client.Usage(ctx, "subscriptions/"+subscriptionID, req, nil)
	if err != nil {
		log.Printf("Error fetching Azure cost data: %v", err)
		return nil, err
	}

	costs := make(map[string]interface{})
	for _, row := range resp.Properties.Rows {
		if len(row) > 1 {
			costs[row[0].(string)] = row[1].(float64)
		}
	}

	if len(costs) == 0 {
		log.Println("No cost data available for the requested period.")
		return nil, errors.New("no cost data available")
	}

	// Store results in Redis cache for 1 hour
	if err := rdb.Set(ctx, cacheKey, costs, time.Hour).Err(); err != nil {
		log.Printf("Error storing data in Redis: %v", err)
	}

	return costs, nil
}
