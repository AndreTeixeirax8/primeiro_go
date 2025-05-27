package validator

import "github.com/go-playground/validator/v10"

func ValidateStruct(obj interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(obj)
	if err != nil {	
		return err}
	return nil
}