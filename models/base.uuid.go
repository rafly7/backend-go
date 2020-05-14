package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	//use gorm for type data varbinary(255)
	//user json for type data varchar(36)
	Id        *uuid.UUID `gorm:"primary_key;type:varchar(36)" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	scope.SetColumn("ID", uuid)
	return nil
}
