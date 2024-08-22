package entities

import "gorm.io/gorm"

type AdminEntity struct {
	gorm.Model
	Xid      string `json:"xid" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Fullname string `json:"fullname" gorm:"column:fullname"`
	Role     string `json:"role" gorm:"column:role"`
}

func (e *AdminEntity) TableName() string {
	return "admin"
}
