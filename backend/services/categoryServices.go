package services

import (
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/models"
	"github.com/man-droid23/e-commerce/repository"
)

type CategoryService interface {
	GetAllCategory() ([]models.Category, error)
	GetCategoryById(id uint) (models.Category, error)
	CreateCategory(category req.CategoryRequest) (models.Category, error)
	UpdateCategory(category req.CategoryRequest, id uint) (models.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) GetAllCategory() ([]models.Category, error) {
	categories, err := s.categoryRepository.GetAllCategory()
	return categories, err
}

func (s *categoryService) GetCategoryById(id uint) (models.Category, error) {
	category, err := s.categoryRepository.GetCategoryById(id)
	return category, err
}

func (s *categoryService) CreateCategory(category req.CategoryRequest) (models.Category, error) {
	categoryModel := models.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	categoryResponse, err := s.categoryRepository.CreateCategory(categoryModel)
	return categoryResponse, err
}

func (s *categoryService) UpdateCategory(category req.CategoryRequest, id uint) (models.Category, error) {
	categoryModel := models.Category{
		Name:        category.Name,
		Description: category.Description,
	}
	categoryResponse, err := s.categoryRepository.UpdateCategory(categoryModel, id)
	return categoryResponse, err
}

func (s *categoryService) DeleteCategory(id uint) error {
	err := s.categoryRepository.DeleteCategory(id)
	return err
}
