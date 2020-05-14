package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDb() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		db_type     = os.Getenv("DB_TYPE")
		db_host     = os.Getenv("DB_HOST")
		db_port     = os.Getenv("DB_PORT")
		db_username = os.Getenv("DB_USERNAME")
		db_password = os.Getenv("DB_PASSWORD")
		db_name     = os.Getenv("DB_NAME")
		configDB    = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)
	)
	db, errDb := gorm.Open(db_type, configDB)
	db.SingularTable(true)
	db.LogMode(true)
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
	// 	&m.User{},
	// 	&m.Permission{},
	// )
	return db, errDb
}
