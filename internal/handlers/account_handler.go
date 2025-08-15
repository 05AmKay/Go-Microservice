package handlers

import (
	"net/http"

	"example.com/api/internal/dto"
	"example.com/api/internal/validation"
	"github.com/gin-gonic/gin"
)

var validate = validation.GetValidator()

func CreateAccount(c *gin.Context) {
	var request dto.CustomerDto

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponseDto(c.Request.URL.Path, http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	validation.ValidateCreateCustomerRequest(c, request)

	c.JSON(http.StatusOK, dto.NewResponseDto(http.StatusOK, "Account created successfully"))
}
