package req

import "github.com/go-playground/validator/v10"

type ProductRequest struct {
	Name        string  `json:"name" validate:"required" form:"name"`
	Description string  `json:"description" validate:"required" form:"description"`
	Price       float64 `json:"price" validate:"required" form:"price"`
	Stock       uint    `json:"stock" validate:"required" form:"stock"`
	Category_Id uint    `json:"category_id" validate:"required" form:"category_id"`
}

func (p *ProductRequest) Validate() error {
	return validator.New().Struct(p)
}
