package repository

import (
	"example.com/api/internal/models"
	"example.com/api/pkg/database"
	"example.com/api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		db: database.GetDatabaseInstance().GetDB(),
	}
}

func (ar *AccountRepository) Save(c *gin.Context, account *models.Account) (*models.Account, error) {
	result := ar.db.Create(account)

	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}

func (ar *AccountRepository) UpdateAccount(c *gin.Context, account *models.Account) (*models.Account, error) {
	result := ar.db.Model(&models.Account{}).Where("account_number = ?", account.AccountNumber).Updates(account)
	if result.Error != nil {
		return nil, result.Error
	}

	// Fetch the updated customer
	account, err := ar.FindByAccountNumber(c, account.AccountNumber)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}

	return account, nil
}

func (ar *AccountRepository) FindByAccountNumber(c *gin.Context, accountNumber string) (*models.Account, error) {
	var account models.Account
	result := ar.db.Where("account_number = ?", accountNumber).First(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No account found
		}
		return nil, result.Error // Error occurred while querying
	}

	return &account, nil
}

func (ar *AccountRepository) FindByCustomerId(c *gin.Context, customerId uint) (*models.Account, error) {
	var account models.Account
	result := ar.db.Where("customer_id = ?", customerId).First(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No account found
		}
		return nil, result.Error // Error occurred while querying
	}

	return &account, nil
}

func (ar *AccountRepository) DeleteAccount(c *gin.Context, account *models.Account) error {
	logger.Sugar.Infof("Deleting account with ID: %+v", account)
	result := ar.db.Delete(account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
