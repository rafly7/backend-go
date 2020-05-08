package services

import (
	"github.com/jinzhu/gorm"
	. "github.com/rafly7/backend/models"
)

type Connection struct {
	Db *gorm.DB
}

type UserService interface {
	GetAllUser() Users
}
