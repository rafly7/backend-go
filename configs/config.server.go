package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConfigServer() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	address := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	return address, os.Getenv("LOG_LEVEL")
}
