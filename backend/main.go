package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/man-droid23/e-commerce/config"
	"github.com/man-droid23/e-commerce/router"
)

func init() {
	config.ConnectDB()
	config.Migrate()
}

func main() {
	defer config.CloseDB()
	app := fiber.New()
	// adding cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.AuthRouter(app)
	router.ProductRouter(app)
	router.CategoryRouter(app)
	log.Fatal(app.Listen(":3000"))
}
