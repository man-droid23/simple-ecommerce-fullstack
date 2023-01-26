package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/controllers"
)

func AuthRouter(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
}
