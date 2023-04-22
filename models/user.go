package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Password is required, minstringlength(6)~Password minimum 6 characters"`
	Age          uint          `gorm:"not null" json:"age" form:"age" valid:"required~Age is required, range(9|100)~Age must be greater than 8"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL" json:"photos"`
	SocialsMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL" json:"socials_media"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL" json:"comments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	return
}
