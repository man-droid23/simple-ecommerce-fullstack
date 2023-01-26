package req

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

func (req *LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}
