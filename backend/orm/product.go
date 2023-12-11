package orm

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Status      string  `gorm:"default:available" json:"status"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
	UserId      uint    `json:"user_id"`
	Username    string  `json:"username"`
}
