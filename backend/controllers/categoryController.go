package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/services"
)

type CategoryController interface {
	GetAllCategory(ctx *fiber.Ctx) error
	GetCategoryById(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategory(ctx *fiber.Ctx) error
	DeleteCategory(ctx *fiber.Ctx) error
}

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *categoryController {
	return &categoryController{categoryService}
}

func (c *categoryController) GetAllCategory(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.GetAllCategory()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get categories",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to get categories",
		"data":    categories,
	})
}

func (c *categoryController) GetCategoryById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	num, errNum := strconv.Atoi(id)
	if errNum != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get category by id",
			"error":   errNum.Error(),
		})
	}
	category, err := c.categoryService.GetCategoryById(uint(num))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get category by id",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to get category by id",
		"data":    category,
	})
}

func (c *categoryController) CreateCategory(ctx *fiber.Ctx) error {
	var category req.CategoryRequest
	err := ctx.BodyParser(&category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
			"error":   err.Error(),
		})
	}
	errValidate := category.Validate()
	if errValidate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
			"error":   errValidate.Error(),
		})
	}
	ctgr, err := c.categoryService.CreateCategory(category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to create category",
		"data":    ctgr,
	})
}

func (c *categoryController) UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	num, errNum := strconv.Atoi(id)
	if errNum != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update category",
			"error":   errNum.Error(),
		})
	}
	var category req.CategoryRequest
	err := ctx.BodyParser(&category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update category",
			"error":   err.Error(),
		})
	}
	errValidate := category.Validate()
	if errValidate != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update category",
			"error":   errValidate.Error(),
		})
	}
	ctgr, err := c.categoryService.UpdateCategory(category, uint(num))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update category",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to update category",
		"data":    ctgr,
	})
}

func (c *categoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	num, errNum := strconv.Atoi(id)
	if errNum != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete category",
			"error":   errNum.Error(),
		})
	}
	err := c.categoryService.DeleteCategory(uint(num))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete category",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to delete category",
	})
}
