package service

import (
	"example.com/api/internal/dto"
	"github.com/gin-gonic/gin"
)

type IAccountService interface {
	CreateAccount(c *gin.Context, customerDto *dto.CustomerDto) error
	UpdateAccount(c *gin.Context, customerDto *dto.CustomerDto) (bool, error)
	FetchAccount(c *gin.Context, mobileNumber string) (*dto.CustomerDto, error)
	DeleteAccount(c *gin.Context, mobileNumber string) error
}
