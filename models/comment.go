package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an Comment
type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"comment_message" form:"comment_message" valid:"required~Comment message is required"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Photo is required"`
	UserID  uint   `gorm:"not null" json:"user_id"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return
	}

	return
}
