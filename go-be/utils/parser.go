package utils

import "github.com/gofiber/fiber/v3"

func BodyParser(c fiber.Ctx, re interface{}) error {
	if err := c.Bind().Body(re); err != nil {
		return err
	}
	return nil
}
