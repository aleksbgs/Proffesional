package controllers

import (
	"github.com/gofiber/fiber/v2"
	"proffesional/models"
	"proffesional/services"
)

// ItineraryController handles HTTP requests for itinerary reconstruction.
type ItineraryController struct {
	service *services.ItineraryService[string]
}

// NewItineraryController creates a new instance of ItineraryController.
func NewItineraryController() *ItineraryController {
	return &ItineraryController{
		service: services.NewItineraryService[string](),
	}
}

// ReconstructItinerary handles POST requests to reconstruct the itinerary.
func (c *ItineraryController) ReconstructItinerary(ctx *fiber.Ctx) error {
	// Parse JSON input into a slice of Ticket arrays for compatibility with challenge input format
	var ticketArrays [][2]string
	if err := ctx.BodyParser(&ticketArrays); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON payload: "+err.Error())
	}

	// Convert [][2]string to []models.Ticket[string]
	tickets := make([]models.Ticket[string], len(ticketArrays))
	for i, arr := range ticketArrays {
		tickets[i] = models.Ticket[string]{Source: arr[0], Destination: arr[1]}
	}

	itinerary, err := c.service.ReconstructItinerary(tickets)
	if err != nil {
		return err
	}

	return ctx.JSON(itinerary)
}
