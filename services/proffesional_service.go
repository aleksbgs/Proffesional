package services

import (
	"github.com/gofiber/fiber/v2"
	"proffesional/models"
)

// ItineraryService handles the business logic for reconstructing itineraries with generics.
type ItineraryService[T comparable] struct{}

// NewItineraryService creates a new instance of ItineraryService.
func NewItineraryService[T comparable]() *ItineraryService[T] {
	return &ItineraryService[T]{}
}

// ReconstructItinerary reconstructs the travel itinerary from a list of tickets.
func (s *ItineraryService[T]) ReconstructItinerary(tickets []models.Ticket[T]) ([]T, error) {
	if len(tickets) == 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "tickets array cannot be empty")
	}

	// Build adjacency map and track incoming edges
	adj := make(map[T]T) // source -> destination
	incoming := make(map[T]bool)
	var zero T // Zero value for type T
	for _, ticket := range tickets {
		src, dest := ticket.Source, ticket.Destination
		if src == zero || dest == zero {
			return nil, fiber.NewError(fiber.StatusBadRequest, "source and destination cannot be empty")
		}
		adj[src] = dest
		incoming[dest] = true
	}

	// Find the starting airport (no incoming edge)
	var start T
	for _, ticket := range tickets {
		if !incoming[ticket.Source] {
			start = ticket.Source
			break
		}
	}
	if start == zero {
		return nil, fiber.NewError(fiber.StatusBadRequest, "no valid starting airport found (possible cycle or invalid itinerary)")
	}

	// Reconstruct the itinerary
	itinerary := []T{start}
	current := start
	visited := make(map[T]bool)

	for len(itinerary) <= len(tickets) {
		next, exists := adj[current]
		if !exists {
			break
		}
		if visited[current] {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid itinerary: cycle detected")
		}
		visited[current] = true
		itinerary = append(itinerary, next)
		current = next
	}

	// Validate the itinerary length
	if len(itinerary) != len(tickets)+1 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid itinerary: incomplete or disconnected path")
	}

	return itinerary, nil
}
