package controllers

import (
	"github.com/kataras/iris/v12"
	. "github.com/rafly7/backend/models"
	. "github.com/rafly7/backend/services"
)

func GetToken(ctx iris.Context, service Connection, user *User) {
	rawJson, err := ctx.GetBody()
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	if data := service.Authenticate(rawJson, user); data != nil {
		ctx.JSON(data)
	} else {
		ctx.StatusCode(404)
	}
}
