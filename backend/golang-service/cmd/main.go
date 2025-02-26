package main

import (
	"log"
	"cloud-cost-explorer/backend/golang-service/config"
	"cloud-cost-explorer/backend/golang-service/internal/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Create new Fiber instance
	app := fiber.New()

	// Register API routes
	api.RegisterRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}