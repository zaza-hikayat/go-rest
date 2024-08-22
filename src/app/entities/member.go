package entities

import (
	"gorm.io/gorm"
	"time"
)

type MemberEntity struct {
	gorm.Model
	RecruiterId *string    `json:"recruiterId"`
	MemberId    string     `json:"memberId" gorm:"unique"`
	FullName    string     `json:"name"`
	Address     string     `json:"address"`
	Password    string     `json:"-"`
	RequiterId  string     `json:"requiterId"`
	Email       string     `json:"email"`
	Gender      string     `json:"gender"`
	Dob         *time.Time `json:"dob" gorm:"column:dob"`
	Phone       string     `json:"phone"`
	Bank        *string    `json:"bank"`
	Rekening    *string    `json:"rekening"`
}

func (u *MemberEntity) TableName() string {
	return "member"
}
