package impl

import (
	"math/rand"
	"strconv"

	"example.com/api/internal/dto"
	errorfactory "example.com/api/internal/error"
	"example.com/api/internal/mapper"
	"example.com/api/internal/models"
	"example.com/api/internal/repository"
	"example.com/api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AccountServiceImpl struct {
}

func generateRandomAccountNumber() string {
	var min = 1000000000
	var max = 9999999999

	// Calculate the range size, ensuring no overflow with uint64.
	// We add 1 to make the range inclusive of the maximum value.
	rangeSize := max - min + 1

	// Generate a random number within the range, then add the minimum.
	// rand.Int63n returns a non-negative pseudo-random 63-bit integer in the half-open interval [0, n).
	// We cast rangeSize to int64 for rand.Int63n and then convert the result to uint64 before adding min.
	randomNumber := strconv.FormatInt(int64(rand.Intn(rangeSize)+min), 10)
	return randomNumber
}

func createNewAccount(savedCustomer *models.Customer, customerDto *dto.CustomerDto) *models.Account {
	newAccount := &models.Account{}
	newAccount.SetCustomerId(savedCustomer.CustomerId)
	newAccount.SetAccountNumber(generateRandomAccountNumber())
	newAccount.SetAccountType(customerDto.AccountsDto.AccountType)
	newAccount.SetBranchAddress("125 Side Street, town")

	return newAccount
}

func (asi *AccountServiceImpl) CreateAccount(c *gin.Context, customerDto *dto.CustomerDto) error {
	logger.Sugar.Infof("Creating account for customer from service: %+v", mapper.MapToCustomerModel(customerDto, &models.Customer{}))
	customer := mapper.MapToCustomerModel(customerDto, &models.Customer{})

	customerRepo := repository.NewCustomerRepository()
	exisitingCustomer, _ := customerRepo.FindByMobileNumber(c, customerDto.Mobilenumber)

	if exisitingCustomer != nil {
		logger.Sugar.Infof("Customer already exists: %+v", exisitingCustomer)
		return errorfactory.CreateCustomerAlreadyExistError(
			c.Request.URL.Path,
			"Customer already registered with given mobileNumber "+customerDto.Mobilenumber,
		)
	}

	savedCustomer, err := customerRepo.Save(c, customer)
	if err != nil {
		logger.Sugar.Errorf("Error saving customer: %v", err)
		return errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}

	newAccount := createNewAccount(savedCustomer, customerDto)

	accountRepo := repository.NewAccountRepository()

	if _, err := accountRepo.Save(c, newAccount); err != nil {
		logger.Sugar.Errorf("Error saving account: %v", err)
		return errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}

	return nil
}

func (asi *AccountServiceImpl) UpdateAccount(c *gin.Context, customerDto *dto.CustomerDto) (bool, error) {
	isUpdated := false
	accountDto := customerDto.GetAccountsDto()
	if accountDto != (&dto.AccountsDto{}) {
		accountRepo := repository.NewAccountRepository()
		existingAccount, err := accountRepo.FindByAccountNumber(c, accountDto.GetAccountNumber())
		if err != nil {
			logger.Sugar.Errorf("Error fetching account by account number: %v", err)
			return false, errorfactory.CreateInternalServerError(c.Request.URL.Path)
		}
		if existingAccount == nil {
			logger.Sugar.Infof("No account found with account number: %s", accountDto.GetAccountNumber())
			return false, errorfactory.CreateResourceNotFoundError(c.Request.URL.Path, "Account", "accountNumber", accountDto.GetAccountNumber())
		}

		account := mapper.MapToAccountsModel(accountDto, existingAccount)
		updatedAccount, err := accountRepo.UpdateAccount(c, account)
		if err != nil {
			logger.Sugar.Errorf("Error updating account: %v", err)
			return false, errorfactory.CreateInternalServerError(c.Request.URL.Path)
		}

		customerRepo := repository.NewCustomerRepository()
		existingCustomer, err := customerRepo.FindByCustomerId(c, updatedAccount.GetCustomerId())
		if err != nil {
			logger.Sugar.Errorf("Error fetching customer by customer id: %v", err)
			return false, errorfactory.CreateInternalServerError(c.Request.URL.Path)
		}
		if existingCustomer == nil {
			logger.Sugar.Infof("No customer found with customer id: %s", updatedAccount.GetCustomerId())
			return false, errorfactory.CreateResourceNotFoundError(c.Request.URL.Path, "Customer", "customerId", updatedAccount.GetCustomerId())
		}

		customer := mapper.MapToCustomerModel(customerDto, existingCustomer)
		_, err = customerRepo.UpdateCustomer(c, customer)
		if err != nil {
			logger.Sugar.Errorf("Error updating customer: %v", err)
			return false, errorfactory.CreateInternalServerError(c.Request.URL.Path)
		}
		isUpdated = true
	}
	return isUpdated, nil
}

func (asi *AccountServiceImpl) FetchAccount(c *gin.Context, mobileNumber string) (*dto.CustomerDto, error) {
	// Implementation logic to fetch account details by mobile number
	// This is a placeholder; actual implementation would involve database calls or external service calls
	customerRepo := repository.NewCustomerRepository()
	customer, err := customerRepo.FindByMobileNumber(c, mobileNumber)

	if err != nil {
		logger.Sugar.Errorf("Error fetching customer by mobile number: %v", err)
		return nil, errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}

	if customer == nil {
		logger.Sugar.Infof("No customer found with mobile number: %s", mobileNumber)
		return nil, errorfactory.CreateResourceNotFoundError(c.Request.URL.Path, "Customer", "mobileNumber", mobileNumber)
	}

	accountRepo := repository.NewAccountRepository()
	account, err := accountRepo.FindByCustomerId(c, customer.CustomerId)
	if err != nil {
		logger.Sugar.Errorf("Error fetching account by customer ID: %v", err)
		return nil, errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}
	if account == nil {
		logger.Sugar.Infof("No account found for customer ID: %d", customer.CustomerId)
		return nil, errorfactory.CreateResourceNotFoundError(c.Request.URL.Path, "Account", "customerId", strconv.Itoa(int(customer.CustomerId)))
	}

	customerDto := mapper.MapToCustomerDto(customer, &dto.CustomerDto{})
	customerDto.SetAccountsDto(*mapper.MapToAccountsDto(account, &customerDto.AccountsDto))

	return customerDto, nil
}

func (asi *AccountServiceImpl) DeleteAccount(c *gin.Context, mobileNumber string) error {
	customerRepo := repository.NewCustomerRepository()
	customer, err := customerRepo.FindByMobileNumber(c, mobileNumber)

	if err != nil {
		logger.Sugar.Errorf("Error fetching customer by mobile number: %v", err)
		return errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}
	if customer == nil {
		logger.Sugar.Infof("No customer found with mobile number: %s", mobileNumber)
		return errorfactory.CreateResourceNotFoundError(c.Request.URL.Path, "Customer", "mobileNumber", mobileNumber)
	}

	accountRepo := repository.NewAccountRepository()
	account, err := accountRepo.FindByCustomerId(c, customer.CustomerId)
	if err != nil {
		logger.Sugar.Errorf("Error fetching account by customer ID: %v", err)
		return errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}
	if account == nil {
		logger.Sugar.Infof("No account found for customer ID: %d", customer.CustomerId)
		return errorfactory.CreateResourceNotFoundError(c.Request.URL.Path, "Account", "customerId", strconv.Itoa(int(customer.CustomerId)))
	}

	err = customerRepo.DeleteCustomer(c, customer)
	if err != nil {
		logger.Sugar.Errorf("Error deleting customer for mobile number: %s: %v", mobileNumber, err)
		return errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}
	err = accountRepo.DeleteAccount(c, account)
	if err != nil {
		logger.Sugar.Errorf("Error deleting account for customer ID %d: %v", customer.CustomerId, err)
		return errorfactory.CreateInternalServerError(c.Request.URL.Path)
	}

	return nil
}
