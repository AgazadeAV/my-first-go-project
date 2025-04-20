package model

// SuccessResponse is a generic response for successful operations.
// swagger:model SuccessResponse
type SuccessResponse struct {
	// A success message
	Message string `json:"message" example:"Operation successful"`
}
