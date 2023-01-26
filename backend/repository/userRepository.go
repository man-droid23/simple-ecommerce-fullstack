package repository

import (
	"github.com/man-droid23/e-commerce/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByEmail(email string) (models.User, error)
	Save(user models.User) (models.User, error)
	Update(user models.User, id uint) (models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	var userResponse models.User
	err := r.db.Create(&user).Scan(&userResponse).Error
	return userResponse, err
}

func (r *userRepository) Update(user models.User, id uint) (models.User, error) {
	var userResponse models.User
	err := r.db.Where("id = ?", id).Save(&user).Scan(&userResponse).Error
	return userResponse, err
}

func (r *userRepository) Delete(id uint) error {
	err := r.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}
