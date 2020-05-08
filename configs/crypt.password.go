package configs

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plaintext string) string {
	byteText := []byte(plaintext)
	hash, err := bcrypt.GenerateFromPassword(byteText, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
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
