package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Todos    []Todo `gorm:"foreignKey:UserID;references:ID"`
}
