package errorx

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Error struct {
	Message  string
	HttpCode int
	Err      error
}

func NewError(httpCode int, message string, err error) *Error {
	return &Error{
		Message:  message,
		HttpCode: httpCode,
		Err:      err,
	}
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func ErrorHandler(c fiber.Ctx, err error) error {
	var inputError *Error
	if errors.As(err, &inputError) {
		return c.Status(inputError.HttpCode).JSON(fiber.Map{
			"error":       inputError.Message,
			"status_code": inputError.HttpCode,
		})

	}
	log.Println(err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":       "Internal Server Error",
		"status_code": fiber.StatusInternalServerError,
	})

}
