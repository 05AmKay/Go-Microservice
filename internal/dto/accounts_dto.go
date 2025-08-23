package dto

type AccountsDto struct {
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type" validate:"required,oneof=Savings Current"`
	BranchAddress string `json:"branch_address"`
}

func (a *AccountsDto) SetAccountNumber(accountNumber string) *AccountsDto {
	a.AccountNumber = accountNumber
	return a
}

// SetAccountType sets the account type
func (a *AccountsDto) SetAccountType(accountType string) *AccountsDto {
	a.AccountType = accountType
	return a
}

// SetBranchAddress sets the branch address
func (a *AccountsDto) SetBranchAddress(branchAddress string) *AccountsDto {
	a.BranchAddress = branchAddress
	return a
}

func (a *AccountsDto) GetAccountNumber() string {
	return a.AccountNumber
}

// GetAccountType returns the account type
func (a *AccountsDto) GetAccountType() string {
	return a.AccountType
}

// GetBranchAddress returns the branch address
func (a *AccountsDto) GetBranchAddress() string {
	return a.BranchAddress
}
