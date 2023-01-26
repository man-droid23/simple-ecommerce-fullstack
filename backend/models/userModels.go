package models

import (
	"time"
)

type User struct {
	User_Id  uint    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Email    string  `json:"email" gorm:"unique;not null;type:varchar(100)"`
	Name     string  `json:"name" gorm:"not null;type:varchar(100)"`
	Password string  `json:"-" gorm:"not null;type:varchar(100)"`
	Address  string  `json:"address" gorm:"not null;type:varchar(100)"`
	Phone    uint64  `json:"phone" gorm:"not null"`
	Role     string  `json:"role" gorm:"default:'user';type:varchar(15);not null"`
	Order    []Order `json:"-" gorm:"foreignKey:User_Id"`
}

type Product struct {
	Product_Id   uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name         string         `json:"name" gorm:"not null;type:varchar(100)" form:"name"`
	Description  string         `json:"description" gorm:"not null;type:varchar(100)" form:"description"`
	Price        float64        `json:"price" gorm:"not null" form:"price"`
	Stock        uint           `json:"stock" gorm:"not null" form:"stock"`
	Category_Id  uint           `json:"category_id" gorm:"not null" form:"category_id"`
	Order_Detail []Order_Detail `json:"-" gorm:"many2many:order_details;foreignKey:Product_Id;references:Product_Id"`
}

type Category struct {
	Category_Id uint      `json:"id" gorm:"primaryKey;autoIncrement;not null" form:"id"`
	Name        string    `json:"name" gorm:"not null;type:varchar(100)" form:"name"`
	Description string    `json:"description" gorm:"not null;type:varchar(100)" form:"description"`
	Product     []Product `json:"-" gorm:"foreignKey:Category_Id"`
}

type Order struct {
	Order_Id     uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	User_Id      uint           `json:"user_id" gorm:"not null"`
	Order_Date   time.Time      `json:"order_date" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Total_Price  float64        `json:"total_price" gorm:"not null"`
	Order_Detail []Order_Detail `json:"-" gorm:"foreignKey:Order_Id"`
}

type Order_Detail struct {
	Order_Detail_Id uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Order_Id        uint `json:"order_id" gorm:"not null"`
	Product_Id      uint `json:"product_id" gorm:"not null"`
	Quantity        uint `json:"quantity" gorm:"not null"`
	Price           uint `json:"price" gorm:"not null"`
}
