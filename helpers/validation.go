package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
)

type ApiError struct {
	Field   string
	Message string
}

func errorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", fieldError.Field())
	case "number":
		return fmt.Sprintf("Field %s must be a number", fieldError.Field())
	case "startswith":
		//tagValue := fieldError.Param()
		return fmt.Sprintf("Field %s must start with %s", fieldError.Field(), fieldError.Param())
	case "min":
		return fmt.Sprintf("Field %s must greater than  %s", fieldError.Field(), fieldError.Param())
	case "max":
		return fmt.Sprintf("Field %s must less than  %s", fieldError.Field(), fieldError.Param())
	case "len":
		return fmt.Sprintf("Field %s must be  %s characters", fieldError.Field(), fieldError.Param())
	case "email":
		return fmt.Sprintf("Field %s must be  a valid email", fieldError.Field())
	}

	return fieldError.Error()
}

func validateImageSize(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(*multipart.FileHeader)
	if !ok {
		return false
	}

	maxFileSize := int64(2 * 1024 * 1024)
	return maxFileSize >= file.Size
}

func ValidateRequest(str interface{}) interface{} {
	validate := validator.New()

	if err := validate.Struct(str); err != nil {
		ve := err.(validator.ValidationErrors)
		errors := make([]ApiError, len(ve))
		for i, fieldError := range ve {
			errors[i] = ApiError{fieldError.Field(), errorMessage(fieldError)}
		}
		return errors
	}
	return nil
}
