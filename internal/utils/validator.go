package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateStruct(s interface{}) validator.ValidationErrors {
	if err := validate.Struct(s); err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			return validateErrs
		}
	}
	return nil
}
