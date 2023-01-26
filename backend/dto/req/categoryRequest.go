package req

import "github.com/go-playground/validator/v10"

type CategoryRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

func (c *CategoryRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
