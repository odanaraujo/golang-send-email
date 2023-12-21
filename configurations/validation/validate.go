package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/odanaraujo/golang/send-emails/configurations/exception"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {

	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		_ = en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateError(validation error) *exception.Exception {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation, &jsonErr) {
		return exception.NewBadRequestErr("invalid field type")
	}
	if errors.As(validation, &jsonValidationError) {
		var errorCauses []exception.Causes
		for _, e := range validation.(validator.ValidationErrors) {
			cause := exception.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorCauses = append(errorCauses, cause)
		}
		return exception.NewRequestValidationError("some fields are invalid", errorCauses)
	}
	return exception.NewBadRequestErr("errir trying to convert fields")
}
