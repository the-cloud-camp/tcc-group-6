package orm

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string
	Status string
	Description string
	Price float64
	Amount float64
}
