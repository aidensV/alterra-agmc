package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"name"`
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}
