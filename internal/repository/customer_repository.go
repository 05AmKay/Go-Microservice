package repository

import (
	"example.com/api/internal/models"
	"example.com/api/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: database.GetDatabaseInstance().GetDB(),
	}
}

func (cr *CustomerRepository) Save(c *gin.Context, customer *models.Customer) (*models.Customer, error) {
	result := cr.db.Create(customer)

	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (cr *CustomerRepository) UpdateCustomer(c *gin.Context, customer *models.Customer) (*models.Customer, error) {
	result := cr.db.Model(&models.Customer{}).Where("customer_id = ?", customer.CustomerId).Updates(customer)
	if result.Error != nil {
		return nil, result.Error
	}
	// Fetch the updated customer
	customer, err := cr.FindByCustomerId(c, customer.CustomerId)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, nil
	}

	return customer, nil
}

func (cr *CustomerRepository) FindByCustomerId(c *gin.Context, customerId uint) (*models.Customer, error) {
	var customer models.Customer
	// Implementation logic to find a customer by mobile number
	// This is a placeholder; actual implementation would involve database calls or external service calls
	result := cr.db.Where("customer_id = ?", customerId).First(&customer)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No customer found
		}
		return nil, result.Error // Error occurred while querying
	}

	return &customer, nil
}

func (cr *CustomerRepository) FindByMobileNumber(c *gin.Context, mobileNumber string) (*models.Customer, error) {
	var customer models.Customer
	// Implementation logic to find a customer by mobile number
	// This is a placeholder; actual implementation would involve database calls or external service calls
	result := cr.db.Where("mobile_number = ?", mobileNumber).First(&customer)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No customer found
		}
		return nil, result.Error // Error occurred while querying
	}

	return &customer, nil
}

func (cr *CustomerRepository) DeleteCustomer(c *gin.Context, customer *models.Customer) error {
	result := cr.db.Delete(customer)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
