package dto

type MobileNumberDto struct {
	MobileNumber string `form:"mobile_number" validate:"required,min=10,max=10"`
}
