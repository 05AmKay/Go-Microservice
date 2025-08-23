package models

type Customer struct {
	Model
	CustomerId   uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null;uniqueIndex:idx_email_unique,where:deleted_at IS NULL"`
	MobileNumber string `gorm:"not null"`
}

// SetName sets the customer name
func (c *Customer) SetName(name string) *Customer {
	c.Name = name
	return c
}

// SetEmail sets the customer email
func (c *Customer) SetEmail(email string) *Customer {
	c.Email = email
	return c
}

// SetMobilenumber sets the customer mobile number
func (c *Customer) SetMobilenumber(mobilenumber string) *Customer {
	c.MobileNumber = mobilenumber
	return c
}

// GetCustomerId returns the customer name
func (c *Customer) GetCustomerId() uint {
	return c.CustomerId
}

// GetName returns the customer name
func (c *Customer) GetName() string {
	return c.Name
}

// GetEmail returns the customer email
func (c *Customer) GetEmail() string {
	return c.Email
}

// GetMobilenumber returns the customer mobile number
func (c *Customer) GetMobilenumber() string {
	return c.MobileNumber
}
