package errs

import "time"

// ErrorResponse represents a structured error response.
// swagger:model ErrorResponse
type ErrorResponse struct {
	// The time at which the error occurred
	Timestamp time.Time `json:"timestamp" example:"2025-04-20T01:23:45Z"`

	// The HTTP status
	Status string `json:"status" example:"400 Bad Request"`

	// A descriptive message explaining the error
	Message string `json:"message" example:"Validation failed"`

	// Detailed field-specific validation errors
	// swagger:type object
	Errors map[string]string `json:"errors,omitempty" example:"email:Email is required,first_name:Invalid format"`
}
