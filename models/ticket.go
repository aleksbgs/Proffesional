package models

// Ticket represents a single flight ticket as a source-destination pair with a generic type.
type Ticket[T comparable] struct {
	Source      T `json:"source"`
	Destination T `json:"destination"`
}
