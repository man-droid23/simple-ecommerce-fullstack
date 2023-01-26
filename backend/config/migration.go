package config

import "github.com/man-droid23/e-commerce/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Order{}, &models.Order_Detail{})
}
