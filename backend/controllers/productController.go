package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/services"
)

type ProductController interface {
	GetAllProduct(ctx *fiber.Ctx) error
	GetProductById(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *productController {
	return &productController{productService}
}

func (c *productController) GetAllProduct(ctx *fiber.Ctx) error {
	products, err := c.productService.GetAllProduct()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed get all product",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get all product",
		"data":    products,
	})
}

func (c *productController) GetProductById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed get product by id",
			"error":   err.Error(),
		})
	}
	product, err := c.productService.GetProductById(uint(num))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed get product by id",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get product by id",
		"data":    product,
	})
}

func (c *productController) CreateProduct(ctx *fiber.Ctx) error {
	var productRequest req.ProductRequest
	err := ctx.BodyParser(&productRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create product",
			"error":   err.Error(),
		})
	}
	validationErr := productRequest.Validate()
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed create product",
			"error":   validationErr.Error(),
		})
	}
	product, err := c.productService.CreateProduct(productRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed create product",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success create product",
		"data":    product,
	})
}

func (c *productController) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed update product",
			"error":   err.Error(),
		})
	}
	var productRequest req.ProductRequest
	err = ctx.BodyParser(&productRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed update product",
			"error":   err.Error(),
		})
	}
	validationErr := productRequest.Validate()
	if validationErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed update product",
			"error":   validationErr.Error(),
		})
	}
	product, err := c.productService.UpdateProduct(productRequest, uint(num))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed update product",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success update product",
		"data":    product,
	})
}

func (c *productController) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed delete product",
			"error":   err.Error(),
		})
	}
	err = c.productService.DeleteProduct(uint(num))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed delete product",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete product",
	})
}
