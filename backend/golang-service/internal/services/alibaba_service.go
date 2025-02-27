package alibaba

import (
	"context"
	"time"
	"log"
	"os"
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/bssopenapi"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: os.Getenv("REDIS_PASSWORD"),
	DB:       0,
})

// FetchAlibabaCloudCosts fetches cost data from Alibaba Cloud Billing API
func FetchAlibabaCloudCosts(startDate, endDate string) (map[string]interface{}, errors) {
	ctx := context.TODO()
	cacheKey := "alibaba-costs-" + startDate + "-" + endDate

	// Check if data is cached
	cachedData, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Returning cached data")
		return map[string]interface{}{"cached": true, "data": cachedData}, nil
	} else if !errors.Is(err, redis.Nil) {
		log.Printf("Error retrieving from cache: %v", err)
	}

	client, err := bssopenapi.NewClientWithAccessKey(
		os.Getenv("ALIBABA_REGION"),
		os.Getenv("ALIBABA_ACCESS_KEY_ID"),
		os.Getenv("ALIBABA_ACCESS_KEY_SECRET"),
	)
	if err != nil {
		log.Printf("Error creating client: %v", err)
		return nil, err
	}

	request := bssopenapi.CreateQueryAccountBalanceRequest()
	request.Scheme = "https"
	request.BillingCycle = startDate[:7]

	response, err := client.QueryAccountBalance(request)
	if err != nil {
		log.Printf("Error querying account balance: %v", err)
		return nil, err
	}

	costs := make(map[string]interface{})
	for _, item := range response.Data.Items.Item {
		costs[item.BillingDate] = item.PretaxAmount
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