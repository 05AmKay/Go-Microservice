package models

type Account struct {
	Model
	AccountNumber string `gorm:"primaryKey;unique;not null"`
	AccountType   string `gorm:"not null"`
	BranchAddress string `gorm:"not null"`
	CustomerId    uint   `gorm:"not null"`
}

func (a *Account) SetAccountNumber(accountNumber string) *Account {
	a.AccountNumber = accountNumber
	return a
}

// SetAccountType sets the account type
func (a *Account) SetAccountType(accountType string) *Account {
	a.AccountType = accountType
	return a
}

// SetBranchAddress sets the branch address
func (a *Account) SetBranchAddress(branchAddress string) *Account {
	a.BranchAddress = branchAddress
	return a
}

// SetCustomerId sets the customer id
func (a *Account) SetCustomerId(customerId uint) *Account {
	a.CustomerId = customerId
	return a
}

func (a *Account) GetAccountNumber() string {
	return a.AccountNumber
}

// GetAccountType returns the account type
func (a *Account) GetAccountType() string {
	return a.AccountType
}

// GetBranchAddress returns the branch address
func (a *Account) GetBranchAddress() string {
	return a.BranchAddress
}

// SetCustomerId sets the customer id
func (a *Account) GetCustomerId() uint {
	return a.CustomerId
}
