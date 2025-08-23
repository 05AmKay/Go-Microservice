package mapper

import (
	"example.com/api/internal/dto"
	"example.com/api/internal/models"
)

func MapToCustomerDto(Customer *models.Customer, CustomerDto *dto.CustomerDto) *dto.CustomerDto {
	return CustomerDto.SetName(Customer.GetName()).
		SetEmail(Customer.GetEmail()).
		SetMobilenumber(Customer.GetMobilenumber())
}

func MapToCustomerModel(CustomerDto *dto.CustomerDto, Customer *models.Customer) *models.Customer {
	return Customer.SetName(CustomerDto.GetName()).
		SetEmail(CustomerDto.GetEmail()).
		SetMobilenumber(CustomerDto.GetMobilenumber())
}
