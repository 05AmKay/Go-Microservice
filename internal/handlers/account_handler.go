package handlers

import (
	"net/http"

	"example.com/api/internal/dto"
	"example.com/api/internal/validation"
	"example.com/api/pkg/logger"
	"github.com/gin-gonic/gin"
)

var validate = validation.GetValidator()

func CreateAccount(c *gin.Context) {
	requestId, _ := c.Get("requestID")
	var request dto.CustomerDto

	logger.Sugar.Infow("CreateAccount called", "requestId", requestId)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseDto(c.Request.URL.Path, http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}

	if err := validation.ValidateCreateCustomerRequest(c, request); err != nil {
		return
	}

	c.JSON(http.StatusOK, dto.NewResponseDto(http.StatusOK, "Account created successfully"))
}
