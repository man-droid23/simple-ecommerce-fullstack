package services

import (
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/models"
	"github.com/man-droid23/e-commerce/repository"
)

type UserService interface {
	FindAll() ([]models.User, error)
	FindByEmail(email string) (models.User, error)
	Save(user req.UserRequest) (models.User, error)
	Update(user req.UserRequest, id uint) (models.User, error)
	Delete(id uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) FindAll() ([]models.User, error) {
	user, err := s.userRepository.FindAll()
	return user, err
}

func (s *userService) FindByEmail(email string) (models.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	return user, err
}

func (s *userService) Save(user req.UserRequest) (models.User, error) {
	var userResponse models.User
	newUser := models.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}
	userResponse, err := s.userRepository.Save(newUser)
	return userResponse, err
}

func (s *userService) Update(user req.UserRequest, id uint) (models.User, error) {
	var userResponse models.User
	newUser := models.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Address:  user.Address,
		Phone:    user.Phone,
	}
	userResponse, err := s.userRepository.Update(newUser, id)
	return userResponse, err
}

func (s *userService) Delete(id uint) error {
	err := s.userRepository.Delete(id)
	return err
}
