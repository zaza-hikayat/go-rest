package entities

import "gorm.io/gorm"

type OrderEntity struct {
	gorm.Model
	OrderNum        string  `gorm:"not null"`
	CreatedBy       string  `gorm:"not null"`
	RecruiterId     *string `gorm:"column:recruiter_id"`
	MemberSnapshot  JSONB   `gorm:"type:json"`
	ProductSnapshot JSONB   `gorm:"type:json"`
	PaymentMethod   string  `gorm:"column:payment_method"`
}

func (e *OrderEntity) TableName() string {
	return "order"
}
