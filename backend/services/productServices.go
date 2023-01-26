package services

import (
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/models"
	"github.com/man-droid23/e-commerce/repository"
)

type ProductService interface {
	GetAllProduct() ([]models.Product, error)
	GetProductById(id uint) (models.Product, error)
	CreateProduct(product req.ProductRequest) (models.Product, error)
	UpdateProduct(product req.ProductRequest, id uint) (models.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository}
}

func (s *productService) GetAllProduct() ([]models.Product, error) {
	products, err := s.productRepository.GetAllProduct()
	return products, err
}

func (s *productService) GetProductById(id uint) (models.Product, error) {
	product, err := s.productRepository.GetProductById(id)
	return product, err
}

func (s *productService) CreateProduct(product req.ProductRequest) (models.Product, error) {
	var productResponse models.Product
	productResponse.Name = product.Name
	productResponse.Description = product.Description
	productResponse.Price = product.Price
	productResponse.Stock = product.Stock
	productResponse.Category_Id = product.Category_Id
	productResponse, err := s.productRepository.CreateProduct(productResponse)
	return productResponse, err
}

func (s *productService) UpdateProduct(product req.ProductRequest, id uint) (models.Product, error) {
	var productResponse models.Product
	productResponse.Name = product.Name
	productResponse.Description = product.Description
	productResponse.Price = product.Price
	productResponse.Stock = product.Stock
	productResponse.Category_Id = product.Category_Id
	productResponse, err := s.productRepository.UpdateProduct(productResponse, id)
	return productResponse, err
}

func (s *productService) DeleteProduct(id uint) error {
	err := s.productRepository.DeleteProduct(id)
	return err
}
