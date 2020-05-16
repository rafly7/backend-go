package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func ConfigServer() (string, string) {
	err := godotenv.Load()
	if err != nil {
		iris.New().Logger().Fatalf("File .env not found for config app server %s", err.Error())
	}
	address := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	return address, os.Getenv("LOG_LEVEL")
}
