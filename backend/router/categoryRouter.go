package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/config"
	"github.com/man-droid23/e-commerce/controllers"
	"github.com/man-droid23/e-commerce/repository"
	"github.com/man-droid23/e-commerce/services"
)

func CategoryRouter(app fiber.Router) {
	categoryController := controllers.NewCategoryController(services.NewCategoryService(repository.NewCategoryRepository(config.DB)))
	category := app.Group("/category")
	category.Get("/", categoryController.GetAllCategory)
	category.Get("/:id", categoryController.GetCategoryById)
	category.Post("/", categoryController.CreateCategory)
	category.Put("/:id", categoryController.UpdateCategory)
	category.Delete("/:id", categoryController.DeleteCategory)
}
