package models

type User struct {
	Base
	Name     string `gorm:"column:name;type:varchar(32);not null" json:"name"`
	Email    string `gorm:"column:email;type:varchar(32);not null" json:"email"`
	Password string `gorm:"column:password;type:varchar(128);not null" json:"password"`
}

type Users []User
