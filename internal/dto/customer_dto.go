package dto

type CustomerDto struct {
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Mobilenumber int    `json:"mobilenumber" validate:"required,gte=18"`
}
