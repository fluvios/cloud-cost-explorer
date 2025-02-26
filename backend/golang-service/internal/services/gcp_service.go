package gcp

import (
	"context"
	"time"
	"log"
	"os"
	"errors"

	"cloud.google.com/go/billing/budgets/apiv1"
	budgetspb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/redis/go-redis/v9"
)

var rdb = redis,NewClient(&redis.Options{
	Addr: "localhost:6379",	
	Password: "",
	DB: 0,
})

// FetchGCPCloudCosts fetches cost data from Google Cloud Billing API
func FetchGCPCloudCosts(startDate, endDate string) (map[string]interface{}, error) {
	ctx := context.TODO()
	cacheKey := "gcp_costs-" + startDate + "-" + endDate

	// Check if data is cached
	cachedData, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {	
		log.Println("Returning cached GCP cost data")
		return map[string]interface{}{"cached": true, "data": cachedData}, nil
	} else if !errors.Is(err, redis.Nil) {
		log.Printf("Error retrieving from cache: %v", err)
	}

	// Create a new budget client
	client, err := budgets.NewBudgetServiceClient(ctx)
	if err != nil {
		log.Printf("Error creating GCP budget client: %v", err)
		return nil, err
	}

	defer client.Close()

	billingAccount := os.Getenv("GCP_BILLING_ACCOUNT")
	if billingAccount == "" {
		log.Println("GCP_BILLING_ACCOUNT environment variable is not set")
		return nil, errors.New("GCP_BILLING_ACCOUNT environment variable is not set")
	}

	req := &budgetspb.ListBudgetsRequest{
		Parent: "billingAccounts/" + billingAccount,
	}

	resp, err := client.ListBudgets(ctx, req)
	if err != nil {
		log.Printf("Error fetching GCP budgets: %v", err)
		return nil, err
	}

	costs := make(map[string]interface{})
	for _, budget := range resp.Budgets {
		costs[budget.Name] = budget.Amount.GetSpecifiedAmount().GetUnits()
	}

	if len(costs) == 0 {
		log.Println("No GCP cost data found")
		return nil, errors.New("No GCP cost data found")
	}

	// Store results in Redis cache for 1 hour
	if err := rdb.Set(ctx, cacheKey, costs, time.Hour).Err(); err != nil {
		log.Printf("Error storing data in Redis: %v", err)
	}

	return costs, nil	
}