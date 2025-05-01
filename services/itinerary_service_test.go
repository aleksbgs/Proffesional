package services

import (
	"github.com/gofiber/fiber/v2"
	"proffesional/models"
	"reflect"
	"testing"
)

// TestItineraryService_ReconstructItinerary tests the ReconstructItinerary method.
func TestItineraryService_ReconstructItinerary(t *testing.T) {
	// Initialize the service with string type for airport codes
	service := NewItineraryService[string]()

	// Define test cases
	tests := []struct {
		name          string
		tickets       []models.Ticket[string]
		expected      []string
		expectedError string
	}{
		{
			name: "Valid itinerary",
			tickets: []models.Ticket[string]{
				{Source: "LAX", Destination: "DXB"},
				{Source: "JFK", Destination: "LAX"},
				{Source: "SFO", Destination: "SJC"},
				{Source: "DXB", Destination: "SFO"},
			},
			expected:      []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
			expectedError: "",
		},
		{
			name:          "Empty tickets array",
			tickets:       []models.Ticket[string]{},
			expected:      nil,
			expectedError: "tickets array cannot be empty",
		},
		{
			name: "Invalid ticket with empty source",
			tickets: []models.Ticket[string]{
				{Source: "", Destination: "DXB"},
				{Source: "JFK", Destination: "LAX"},
			},
			expected:      nil,
			expectedError: "source and destination cannot be empty",
		},
		{
			name: "Invalid ticket with empty destination",
			tickets: []models.Ticket[string]{
				{Source: "LAX", Destination: ""},
				{Source: "JFK", Destination: "LAX"},
			},
			expected:      nil,
			expectedError: "source and destination cannot be empty",
		},
		{
			name: "Disconnected itinerary",
			tickets: []models.Ticket[string]{
				{Source: "JFK", Destination: "LAX"},
				{Source: "SFO", Destination: "SJC"},
			},
			expected:      nil,
			expectedError: "invalid itinerary: incomplete or disconnected path",
		},
		{
			name: "Single ticket",
			tickets: []models.Ticket[string]{
				{Source: "JFK", Destination: "LAX"},
			},
			expected:      []string{"JFK", "LAX"},
			expectedError: "",
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the method
			result, err := service.ReconstructItinerary(tt.tickets)

			// Check error
			if tt.expectedError == "" {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q, got nil", tt.expectedError)
				} else if fiberErr, ok := err.(*fiber.Error); ok {
					if fiberErr.Message != tt.expectedError {
						t.Errorf("expected error %q, got %q", tt.expectedError, fiberErr.Message)
					}
				} else {
					t.Errorf("expected fiber.Error, got %T", err)
				}
			}

			// Check result
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected itinerary %v, got %v", tt.expected, result)
			}
		})
	}
}
