package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// . "github.com/rafly7/backend/models"
)

type CreditCard struct {
	gorm.Model
	Number string
	UID    string
}

type User struct {
	gorm.Model
	Name       string     `sql:"index"`
	CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/simple?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DropTableIfExists(User{}, CreditCard{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&User{},
		&CreditCard{},
	)
	var card CreditCard
	db.Model(&User{}).Related(&card, "CreditCard")
	//db.Model(&User{}).AddForeignKey("permission_id", "permission(id)", "RESTRICT", "RESTRICT")
	// Migrate(db, users, permissions)
}
