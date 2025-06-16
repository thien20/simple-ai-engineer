package utils

import "github.com/gofiber/fiber/v3"

func BodyParser(c fiber.Ctx, request interface{}) error {
	if err := c.Bind().Body(request); err != nil {
		return err
	}
	return nil
}
