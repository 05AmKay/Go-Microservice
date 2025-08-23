package mapper

import (
	"example.com/api/internal/dto"
	"example.com/api/internal/models"
)

func MapToAccountsDto(Accounts *models.Account, AccountsDto *dto.AccountsDto) *dto.AccountsDto {
	return AccountsDto.SetAccountNumber(Accounts.AccountNumber).
		SetAccountType(Accounts.AccountType).
		SetBranchAddress(Accounts.BranchAddress)
}

func MapToAccountsModel(AccountsDto *dto.AccountsDto, Accounts *models.Account) *models.Account {
	return Accounts.SetAccountNumber(AccountsDto.GetAccountNumber()).
		SetAccountType(AccountsDto.GetAccountType()).
		SetBranchAddress(AccountsDto.GetBranchAddress())

}
