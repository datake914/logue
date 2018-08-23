package infrastructure

import (
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

// CustomValidator represents the validator.
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator returns a pointer to the CustomValidator struct.
func NewCustomValidator() *CustomValidator {
	v := &CustomValidator{
		validator: validator.New(),
	}
	// Register tag name function to get json tag filed name.
	v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return v
}

// Validate validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
