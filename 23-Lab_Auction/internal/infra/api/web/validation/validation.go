package validation

import (
	"encoding/json"
	"errors"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate  = validator.New()
	translate ut.Translator
)

func init() {
	value, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		en := en.New()
		enTranslation := ut.New(en, en)

		translate, _ = enTranslation.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, translate)
	}
}

func ValidateErr(validationError error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid type error")
	}

	if errors.As(validationError, &jsonValidation) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(translate),
			})
		}

		return rest_err.NewBadRequestError("Invalid field values", errorCauses...)
	}

	return rest_err.NewBadRequestError("Error tring to convert fields")
}
