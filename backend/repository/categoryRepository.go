package repository

import (
	"github.com/man-droid23/e-commerce/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory() ([]models.Category, error)
	GetCategoryById(id uint) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category, id uint) (models.Category, error)
	DeleteCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryById(id uint) (models.Category, error) {
	var category models.Category
	err := r.db.Where("category_id = ?", id).Find(&category).Error
	return category, err
}

func (r *categoryRepository) CreateCategory(category models.Category) (models.Category, error) {
	var categoryResponse models.Category
	err := r.db.Create(&category).Scan(&categoryResponse).Error
	return categoryResponse, err
}

func (r *categoryRepository) UpdateCategory(category models.Category, id uint) (models.Category, error) {
	var categoryResponse models.Category
	err := r.db.Where("category_id = ?", id).Updates(&category).Scan(&categoryResponse).Error
	return categoryResponse, err
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	err := r.db.Where("category_id = ?", id).Delete(&models.Category{}).Error
	return err
}
