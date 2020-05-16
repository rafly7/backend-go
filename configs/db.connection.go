package configs

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func ConnectToDb() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		iris.New().Logger().Fatalf("File .env not found for config database %s", err.Error())
	}
	var (
		db_type     = os.Getenv("DB_TYPE")
		db_host     = os.Getenv("DB_HOST")
		db_port     = os.Getenv("DB_PORT")
		db_username = os.Getenv("DB_USERNAME")
		db_password = os.Getenv("DB_PASSWORD")
		db_name     = os.Getenv("DB_NAME")
		db_log      = os.Getenv("DB_LOG_LEVEL")
		configDB    = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)
	)
	db, errDb := gorm.Open(db_type, configDB)
	db.SingularTable(true)
	if db_log == "true" {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}
	// db.LogMode(true)
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
	// 	&m.User{},
	// 	&m.Permission{},
	// )
	return db, errDb
}
