package repository

import (
	"github.com/man-droid23/e-commerce/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]models.Product, error)
	GetProductById(id uint) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product, id uint) (models.Product, error)
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductById(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).Find(&product).Error
	return product, err
}

func (r *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	var productResponse models.Product
	err := r.db.Create(&product).Scan(&productResponse).Error
	return productResponse, err
}

func (r *productRepository) UpdateProduct(product models.Product, id uint) (models.Product, error) {
	var productResponse models.Product
	err := r.db.Where("id = ?", id).Updates(&product).Scan(&productResponse).Error
	return productResponse, err
}

func (r *productRepository) DeleteProduct(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&models.Product{}).Error
	return err
}
