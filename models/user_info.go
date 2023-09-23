package models

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	*gorm.Model
	Email    string
	Password string
}
