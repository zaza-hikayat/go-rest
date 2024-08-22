package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	app_errors "github.com/zaza-hikayat/go-rest-sample/src/app/errors"
)

type errorResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

func NewErrorHandler() func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if err, ok := err.(*app_errors.CommonError); ok {
			if err.HttpCode > 0 {
				code = err.HttpCode
			}

			return ctx.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": err.ErrorMessage,
			})
		}

		if errs, ok := err.(validator.ValidationErrors); ok {
			var validationErrors []errorResponse
			for _, err := range errs {
				elem := errorResponse{err.StructField(), msgForTag(err)}
				validationErrors = append(validationErrors, elem)
			}

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors":  validationErrors,
				"message": "validation errors",
			})
		}

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		if code == fiber.StatusNotFound {
			return ctx.Status(code).JSON(fiber.Map{
				"error":   true,
				"message": "Resource not found",
			})
		}

		return ctx.Status(code).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "number":
		return "Invalid number"
	}
	return fe.Error() // default error
}
