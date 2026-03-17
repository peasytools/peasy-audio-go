package peasyaudio

import "fmt"

// PeasyError represents an API error response.
type PeasyError struct {
	StatusCode int
	Message    string
}

func (e *PeasyError) Error() string {
	return fmt.Sprintf("peasyaudio: api error (HTTP %d): %s", e.StatusCode, e.Message)
}

// NotFoundError is returned when a resource is not found (404).
type NotFoundError struct {
	Resource   string
	Identifier string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("peasyaudio: %s not found: %s", e.Resource, e.Identifier)
}
