package user

type CreateUserInput struct {
	Name string `json:"name" example:"Azer"`
	Age  int    `json:"age" example:"30"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"Internal server error"`
}
