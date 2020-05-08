package models

type User struct {
	Base
	Name     string `gorm:"column:name;type:varchar(32);not null"`
	Email    string `gorm:"column:email;type:varchar(32);not null"`
	Password string `gorm:"column:password;type:varchar(128);not null"`
}

type Users []User
