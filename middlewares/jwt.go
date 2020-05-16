package middlewares

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joho/godotenv"
)

func JwtHandler() *jwtmiddleware.Middleware {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
		ContextKey:    "token",
		SigningMethod: jwt.SigningMethodHS256,
	})
}
