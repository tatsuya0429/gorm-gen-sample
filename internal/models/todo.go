package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID      int64
	Title       string
	Completed   bool
	Description string
	User        User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
