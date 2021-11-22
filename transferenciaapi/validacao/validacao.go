package validacao

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func GetInstance() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}
