package handlers

import (
	"fmt"
	"net/http"

	"example.com/api/internal/dto"
	errorfactory "example.com/api/internal/error"
	exceptions "example.com/api/internal/exception"
	"example.com/api/internal/service"
	"example.com/api/internal/validation"
	"example.com/api/pkg/logger"
	"github.com/gin-gonic/gin"
)

// AccountHandler handles account-related requests
type AccountHandler struct {
	accountService service.IAccountService
}

func NewAccountHandler(accountService service.IAccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (ah *AccountHandler) CreateAccount(c *gin.Context) {
	fmt.Printf("Account service instance %v", ah.accountService)
	requestId, _ := c.Get("requestID")
	var customerDto dto.CustomerDto

	logger.Sugar.Infow("CreateAccount called", "requestId", requestId)

	if err := c.ShouldBindJSON(&customerDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseDto(c.Request.URL.Path, http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}

	if err := validation.ValidateCreateCustomerRequest(c, customerDto); err != nil {
		return
	}

	logger.Sugar.Infof("CreateAccount: Validated customer data: %+v", customerDto)

	err := ah.accountService.CreateAccount(c, &customerDto)
	if err != nil {
		if appErr, ok := err.(errorfactory.ApplicationError); ok {
			exceptions.ThrowError(c, appErr)
		}
		logger.Sugar.Errorf("CreateAccount: Error creating account: %v", err)
		return
	}

	c.JSON(http.StatusOK, dto.NewResponseDto(http.StatusOK, "Account created successfully", nil))
}

func (ah *AccountHandler) FetchAccountDetails(c *gin.Context) {
	var mobileNumberDto dto.MobileNumberDto
	requestId, _ := c.Get("requestID")
	logger.Sugar.Infow("FetchAccountDetails called", "requestId", requestId)

	if err := c.ShouldBindQuery(&mobileNumberDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseDto(c.Request.URL.Path, http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}

	logger.Sugar.Infof("FetchAccount: Validating customer data: %+v", mobileNumberDto)

	if err := validation.ValidateCreateCustomerRequest(c, mobileNumberDto); err != nil {
		return
	}

	customerDto, err := ah.accountService.FetchAccount(c, mobileNumberDto.MobileNumber)
	if err != nil {
		if appErr, ok := err.(errorfactory.ApplicationError); ok {
			exceptions.ThrowError(c, appErr)
		}
		logger.Sugar.Errorf("FetchAccount: Error Fetching account details: %v", err)
		return
	}

	c.JSON(http.StatusOK, dto.NewResponseDto(http.StatusOK, "Account details fetched successfully", customerDto))
}

func (ah *AccountHandler) UpdateAccountDetails(c *gin.Context) {
	var requestId, _ = c.Get("requestID")
	logger.Sugar.Infow("UpdateAccountDetails called", "requestId", requestId)

	var customerDto dto.CustomerDto
	if err := c.ShouldBindJSON(&customerDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseDto(c.Request.URL.Path, http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}

	if err := validation.ValidateCreateCustomerRequest(c, customerDto); err != nil {
		return
	}

	isUpdated, err := ah.accountService.UpdateAccount(c, &customerDto)
	if err != nil {
		if appErr, ok := err.(errorfactory.ApplicationError); ok {
			exceptions.ThrowError(c, appErr)
		}
		logger.Sugar.Errorf("UpdateAccountDetails: Error updating account details: %v", err)
		return
	}
	if !isUpdated {
		c.JSON(http.StatusExpectationFailed, dto.NewErrorResponseDto(c.Request.URL.Path, "Update operation failed. Please try again or contact Dev team", http.StatusExpectationFailed))
	} else {
		c.JSON(http.StatusOK, dto.NewResponseDto(http.StatusOK, "Account updated successfully", nil))
	}
}

func (ah *AccountHandler) DeleteAccount(c *gin.Context) {
	var mobileNumberDto dto.MobileNumberDto
	requestId, _ := c.Get("requestID")
	logger.Sugar.Infow("DeleteAccount called", "requestId", requestId)
	if err := c.ShouldBindQuery(&mobileNumberDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseDto(c.Request.URL.Path, http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}

	if err := validation.ValidateCreateCustomerRequest(c, mobileNumberDto); err != nil {
		return
	}

	err := ah.accountService.DeleteAccount(c, mobileNumberDto.MobileNumber)
	if err != nil {
		if appErr, ok := err.(errorfactory.ApplicationError); ok {
			exceptions.ThrowError(c, appErr)
		}
		logger.Sugar.Errorf("DeleteAccount: Error deleting account: %v", err)
		return
	}
	c.JSON(http.StatusOK, dto.NewResponseDto(http.StatusOK, "Account deleted successfully", nil))
}
