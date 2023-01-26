package req

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Name     string `json:"name" form:"name" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	Address  string `json:"address" form:"address" validate:"required"`
	Phone    uint64 `json:"phone" form:"phone" validate:"required"`
}

func (req *UserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}
