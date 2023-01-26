package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/config"
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/repository"
	"github.com/man-droid23/e-commerce/services"
)

func Login(c *fiber.Ctx) error {
	reqLogin := new(req.LoginRequest)
	if err := c.BodyParser(reqLogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	errValidate := reqLogin.Validate()
	if errValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}
	userRepository := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthServices(userService)
	token, err := authService.Login(*reqLogin)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	})
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
	})
}

func Register(c *fiber.Ctx) error {
	reqRegister := new(req.UserRequest)
	if err := c.BodyParser(reqRegister); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	errValidate := reqRegister.Validate()
	if errValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}
	userRepository := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthServices(userService)
	token, err := authService.Register(*reqRegister)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": token,
	})
}
