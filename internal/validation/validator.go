package validation

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validatorObj *validator.Validate
	once         sync.Once
)

// InitValidator initializes the validator instance thread-safely
func InitValidator() {
	once.Do(func() {
		validatorObj = validator.New()
	})
}

// GetValidator returns the validator instance thread-safely
func GetValidator() *validator.Validate {
	if validatorObj == nil {
		InitValidator()
	}
	return validatorObj
}
