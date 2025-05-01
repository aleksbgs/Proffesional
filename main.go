package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"proffesional/controllers"
)

func main() {
	// Initialize Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Initialize controller
	itineraryController := controllers.NewItineraryController()

	// Routes
	app.Post("api/itinerary", itineraryController.ReconstructItinerary)

	// Start server
	app.Listen(":8080")
}
