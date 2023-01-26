package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	validateToken, err := utils.ValidateToken(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	c.Locals("token", validateToken)
	return c.Next()
}
