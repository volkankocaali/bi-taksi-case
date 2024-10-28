package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/volkankocaali/bi-taksi-case/internal/resource/response"
	"reflect"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) getJSONFieldName(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return field.Name
	}
	return jsonTag
}

func (v *Validator) ValidateStruct(s interface{}) []response.ValidationError {
	var errors []response.ValidationError

	err := v.validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(s).Elem().FieldByName(err.Field())
			jsonFieldName := v.getJSONFieldName(field)

			errors = append(errors, response.ValidationError{
				Field: jsonFieldName,
				Error: fmt.Sprintf("failed on '%s' condition", err.Tag()),
				Param: err.Param(),
			})
		}
	}

	return errors
}
