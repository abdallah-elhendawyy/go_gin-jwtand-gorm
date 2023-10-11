package mosels

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID uint
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}
