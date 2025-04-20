package model

import "github.com/google/uuid"

// CreateUserInput represents the body for creating a new user.
// swagger:model CreateUserInput
type CreateUserInput struct {
	// First name of the user. Must be Latin, 1-50 characters.
	FirstName *string `json:"first_name" example:"Azer"`

	// Last name of the user. Must be Latin, 1-50 characters.
	LastName *string `json:"last_name" example:"Agazade"`

	// Unique username (Latin letters, numbers, dash or underscore).
	Username *string `json:"username" example:"azer_agazade"`

	// Email address in valid format.
	Email *string `json:"email" example:"azer.agazade05@yandex.ru"`

	// Phone number starting with +7 and 10 digits.
	PhoneNumber *string `json:"phone_number" example:"+79672417678"`

	// Birthdate in format YYYY-MM-DD.
	BirthDate *string `json:"birth_date" example:"1995-05-03"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	BirthDate   string    `json:"birth_date"`
}
