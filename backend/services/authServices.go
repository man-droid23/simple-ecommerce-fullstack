package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/man-droid23/e-commerce/dto/req"
	"github.com/man-droid23/e-commerce/utils"
)

type AuthServices interface {
	Login(userLogin req.LoginRequest) (string, error)
	Register(userRegister req.UserRequest) (string, error)
	Validate(token string) (*jwt.Token, error)
}

type authServices struct {
	userService UserService
}

func NewAuthServices(userService UserService) *authServices {
	return &authServices{userService}
}

func (s *authServices) Login(userLogin req.LoginRequest) (string, error) {
	user, err := s.userService.FindByEmail(userLogin.Email)
	if err != nil {
		return "", err
	}
	errPass := utils.CheckPasswordHash(userLogin.Password, user.Password)
	if errPass != nil {
		return "", errPass
	}
	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token, errToken := utils.GenerateJWTToken(&claims)
	if errToken != nil {
		return "", errToken
	}
	return token, nil
}

func (s *authServices) Register(userRegister req.UserRequest) (string, error) {
	newUser := req.UserRequest{
		Email:    userRegister.Email,
		Name:     userRegister.Name,
		Password: userRegister.Password,
		Address:  userRegister.Address,
		Phone:    userRegister.Phone,
	}
	hash, errHash := utils.HashPassword(newUser.Password)
	if errHash != nil {
		return "", errHash
	}
	newUser.Password = hash
	_, err := s.userService.Save(newUser)
	if err != nil {
		return "", err
	}
	return "Register Success", nil
}

func (s *authServices) Validate(token string) (*jwt.Token, error) {
	validateToken, err := utils.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	return validateToken, nil
}
