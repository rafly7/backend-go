package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/rafly7/backend/configs"
	// . "github.com/rafly7/backend/models"
)

type Base struct {
	//use gorm for type data varbinary(255)
	//user json for type data varchar(36)
	Id        *uuid.UUID `gorm:"primary_key;type:varchar(36)"`
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

type Permission struct {
	Base
	PermissionName string `gorm:"type:varchar(8);not null"`
}

type Permissions []Permission

type User struct {
	Base
	Name     string `gorm:"column:name;type:varchar(32);not null"`
	Email    string `gorm:"column:email;type:varchar(32);not null"`
	Password string `gorm:"column:password;type:varchar(128);not null"`
	// PermissionId Permission `gorm:"foreignkey:permissionId;association_foreignkey:permission(id);column:permissionId"`
	PermissionId Permission `gorm:"association_foreignkey:id"`
}

type Users []User

var (
	users = Users{
		User{Name: "Agna", Email: "agna@gmail.com", Password: configs.HashPassword("123456")},
		User{Name: "Rafly", Email: "rafly@gmail.com", Password: configs.HashPassword("dn38d4n4")},
		User{Name: "Vito", Email: "vito@gmail.com", Password: configs.HashPassword("jdabswbd")},
	}
	permissions = Permissions{
		Permission{PermissionName: "admin"},
		Permission{PermissionName: "user"},
	}
)

func main() {
	db, err := configs.ConnectToDb()
	if err != nil {
		panic(err.Error())
	}
	db.DropTableIfExists(User{}, Permission{})
	db.AutoMigrate(User{}, Permission{})
	db.Model(&User{}).Related(&Permission{}, "PermissionId")
	//db.Model(&User{}).AddForeignKey("permission_id", "permission(id)", "RESTRICT", "RESTRICT")
	// Migrate(db, users, permissions)
}

func Migrate(db *gorm.DB, models ...interface{}) {

	for _, values := range *&models {
		items1, err1 := values.(Users)
		items2, err2 := values.(Permissions)
		if err1 == true {
			for _, i := range items1 {
				db.NewRecord(i)
				db.Create(&i)
				db.NewRecord(i)
			}
		}
		if err2 == true {
			for _, i := range items2 {
				db.NewRecord(i)
				db.Create(&i)
				db.NewRecord(i)
			}
		}
	}
}
