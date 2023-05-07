package router

import (
	"_web/pkg"

	"github.com/go-playground/validator/v10"
)

type UnauthorizedResponse struct {
	Message string `json:"message"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message"`
}

type NotFoundResponse struct {
	Message string `json:"message"`
}

type BadRequestResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type BadRequestResponseWithErrors struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func BadRequestFromError(err error) BadRequestResponseWithErrors {
	var errors = make(map[string]string)

	if _, ok := err.(*validator.InvalidValidationError); ok {
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Tag()] = pkg.ParseErrorMessage(pkg.ErrorMessages[err.Tag()], err.StructField())
	}

	return BadRequestResponseWithErrors{
		Message: "Bad Request",
		Errors:  errors,
	}
}
