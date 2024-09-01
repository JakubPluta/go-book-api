package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

const alphaSpaceRegexString string = "^[a-zA-Z ]+$"

// New returns a new validator.Validate instance that is configured to
// recognize struct tags of the form "json", and to validate the
// "alphaspace" tag. The value of the "json" tag is used as the name of the
// field in error messages. If the value is '-', the field is ignored.
func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("alphaspace", isAlphaSpace)

	return validate
}

// isAlphaSpace is a validation function for the "alphaspace" tag. It returns
// true if the field's value consists only of alphabetic characters and
// whitespace, and false otherwise.
func isAlphaSpace(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(alphaSpaceRegexString)
	return reg.MatchString(fl.Field().String())
}
