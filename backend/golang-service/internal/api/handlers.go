package api

import (
	"github.com/gofiber/fiber/v2"
	"cloud-cost-explorer/backend/golang-service/internal/services/aws"
	"cloud-cost-explorer/backend/golang-service/internal/services/gcp"
	"cloud-cost-explorer/backend/golang-service/internal/services/azure"
	"cloud-cost-explorer/backend/golang-service/internal/services/alibaba"
)

// RegisterRoutes sets up API routes
func RegisterRoutes(app *fiber.App) {
	app.Get("/health", HealthCheck)
	app.Get("/costs/aws", GetAWSCloudCosts)
	app.Get("/costs/gcp", GetGCPCloudCosts)
	app.Get("/costs/azure", GetAzureCloudCosts)
	app.Get("/costs/alibaba", GetAlibabaCloudCosts)
}

// validateDateFormat checks if the date is in YYYY-MM-DD format
func validateDateFormat(date string) error {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(date) {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}
	return nil
}

// HealthCheck returns the health status of the service
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

// GetAWSCloudCosts fetches AWS cost data with dynamic date range
func GetAWSCloudCosts(c *fiber.Ctx) error {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		return c.Status(400).JSON(fiber.Map{"error": "start_date and end_date query parameters are required"})
	}

	if err := validateDateFormat(startDate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid start_date format, expected YYYY-MM-DD"})
	}

	if err := validateDateFormat(endDate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid end_date format, expected YYYY-MM-DD"})
	}

	costs, err := aws.FetchAWSCloudCosts(startDate, endDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(costs)
}

// GetGCPCloudCosts fetches GCP cost data
func GetGCPCloudCosts(c *fiber.Ctx) error {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		return c.Status(400).JSON(fiber.Map{"error": "start_date and end_date query parameters are required"})
	}

	if err := validateDateFormat(startDate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid start_date format, expected YYYY-MM-DD"})
	}

	if err := validateDateFormat(endDate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid end_date format, expected YYYY-MM-DD"})
	}

	costs, err := gcp.FetchGCPCloudCosts(startDate, endDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(costs)
}

// GetAzureCloudCosts fetches Azure cost data
func GetAzureCloudCosts(c *fiber.Ctx) error {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		return c.Status(400).JSON(fiber.Map{"error": "start_date and end_date query parameters are required"})
	}

	if err := validateDateFormat(startDate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid start_date format, expected YYYY-MM-DD"})
	}

	if err := validateDateFormat(endDate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid end_date format, expected YYYY-MM-DD"})
	}

	costs, err := azure.FetchAzureCloudCosts(startDate, endDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(costs)
}

// GetAlibabaCloudCosts fetches Alibaba cost data
func GetAlibabaCloudCosts(c *fiber.Ctx) error {
	costs, err := alibaba.FetchAlibabaCloudCosts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(costs)
}
