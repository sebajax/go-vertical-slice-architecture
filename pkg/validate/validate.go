package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// generate validator instance
var Validator = validator.New(validator.WithRequiredStructEnabled())

func Validate(body interface{}) (string, error) {
	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Sprintf("Field %s Tag %s", err.Field(), err.Tag()), err
		}
	}
	return "", nil
}
