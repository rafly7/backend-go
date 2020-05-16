package services

import (
	"encoding/json"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/rafly7/backend/configs"
	"github.com/rafly7/backend/models"
)

func (c *Connection) Authenticate(payload []byte, u *models.User) map[string]string {
	if confErr := godotenv.Load(); confErr != nil {
		return nil
	}
	if err := json.Unmarshal(payload, &u); err != nil {
		return nil
	}
	password := u.Password
	if exists := c.Db.Where("email = ?", u.Email).First(&u).RecordNotFound(); exists {
		return nil
	} else {
		const expires int64 = 60 * 60 * 24
		if verify := configs.VerifyPassword(u.Password, password); verify {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"name":  u.Name,
				"email": u.Email,
			})
			tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
			if errToken != nil {
				return nil
			}
			return map[string]string{
				"token": tokenString,
			}
		} else {
			iris.New().Logger().Errorf("Validation Password: %s", password)
			return nil
		}
	}
}
