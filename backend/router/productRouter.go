package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/config"
	"github.com/man-droid23/e-commerce/controllers"
	"github.com/man-droid23/e-commerce/repository"
	"github.com/man-droid23/e-commerce/services"
)

func ProductRouter(app *fiber.App) {
	productController := controllers.NewProductController(services.NewProductService(repository.NewProductRepository(config.DB)))
	product := app.Group("/product")
	product.Get("/", productController.GetAllProduct)
	product.Get("/:id", productController.GetProductById)
	product.Post("/", productController.CreateProduct)
	product.Put("/:id", productController.UpdateProduct)
	product.Delete("/:id", productController.DeleteProduct)
}
