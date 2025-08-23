package dto

type CustomerDto struct {
	Name         string      `json:"name" validate:"required,min=5,max=30"`
	Email        string      `json:"email" validate:"required,email"`
	Mobilenumber string      `json:"mobile_number" form:"mobile_number" validate:"required,min=10,max=10"`
	AccountsDto  AccountsDto `json:"accounts" validate:"required"`
}

// SetName sets the customer name
func (c *CustomerDto) SetName(name string) *CustomerDto {
	c.Name = name
	return c
}

// SetEmail sets the customer email
func (c *CustomerDto) SetEmail(email string) *CustomerDto {
	c.Email = email
	return c
}

// SetMobilenumber sets the customer mobile number
func (c *CustomerDto) SetMobilenumber(mobilenumber string) *CustomerDto {
	c.Mobilenumber = mobilenumber
	return c
}

// SetAccountsDto sets the accounts details
func (c *CustomerDto) SetAccountsDto(accountsDto AccountsDto) *CustomerDto {
	c.AccountsDto = accountsDto
	return c
}

// GetName returns the customer name
func (c *CustomerDto) GetName() string {
	return c.Name
}

// GetEmail returns the customer email
func (c *CustomerDto) GetEmail() string {
	return c.Email
}

// GetMobilenumber returns the customer mobile number
func (c *CustomerDto) GetMobilenumber() string {
	return c.Mobilenumber
}

// GetAccountsDto returns the accounts details
func (c *CustomerDto) GetAccountsDto() *AccountsDto {
	return &c.AccountsDto
}
