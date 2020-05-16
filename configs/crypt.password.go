package configs

import (
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plaintext string) string {
	byteText := []byte(plaintext)
	hash, err := bcrypt.GenerateFromPassword(byteText, bcrypt.MinCost)
	if err != nil {
		iris.New().Logger().Errorf("Error Hash Password: %s", err.Error())
	}
	return string(hash)
}

func VerifyPassword(hashpwd, plainText string) bool {
	byteHash, byteText := []byte(hashpwd), []byte(plainText)
	err := bcrypt.CompareHashAndPassword(byteHash, byteText)
	if err != nil {
		//panic(err.Error())
		return false
	}
	return true
}
