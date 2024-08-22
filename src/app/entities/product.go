package entities

import "gorm.io/gorm"

type ProductEntity struct {
	gorm.Model
	Code        string `gorm:"not null,unique"`
	Name        string `gorm:"varchar(255)"`
	Qty         uint   `gorm:"default:0"`
	Description string `gorm:"text"`
	Image       string `gorm:"text"`
}

func (e *ProductEntity) TableName() string {
	return "product"
}
