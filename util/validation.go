package util

import "github.com/go-playground/validator"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct performs struct field validation using the validator package.
func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}
