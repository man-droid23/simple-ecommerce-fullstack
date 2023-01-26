package config

import "github.com/man-droid23/e-commerce/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.Order_Detail{})
}
