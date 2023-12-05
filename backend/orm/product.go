package orm

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Status      string `gorm:"default:available"`
	Description string
	Price       float64
	Amount      float64
	UserId      uint
	Username    string
}
