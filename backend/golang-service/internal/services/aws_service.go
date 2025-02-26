package aws

import (
	"context"
	"time"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

// FetchAWSCloudCosts fetches cost data from AWS Cost Explorer
func FetchAWSCloudCosts() (map[string]interface{}, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Printf("Error loading AWS config: %v", err)
		return nil, err
	}

	client := costexplorer.NewFromConfig(cfg)

	// Define time period (last 7 days)
	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(startDate),
			End:   aws.String(endDate),
		},
		Granularity: aws.String("DAILY"),
		Metrics:     []string{"BlendedCost"},
	}

	result, err := client.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		log.Printf("Error fetching AWS cost data: %v", err)
		return nil, err
	}

	costs := make(map[string]interface{})
	for _, res := range result.ResultsByTime {
		date := *res.TimePeriod.Start
		costs[date] = res.Total["BlendedCost"].Amount
	}

	return costs, nil
}
