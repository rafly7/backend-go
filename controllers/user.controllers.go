package controllers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	. "github.com/rafly7/backend/models"
	. "github.com/rafly7/backend/services"
)

func GetUser(ctx iris.Context, service Connection, user *User, users *Users) {
	id := ctx.URLParam("id")
	if len(id) == 0 {
		fmt.Println("GET ALL USER")
		if data := service.GetAllUser(users); data != nil {
			ctx.JSON(data)
		} else {
			ctx.StatusCode(404)
			iris.New().Logger().Errorf("%s", "ERROR NOT FOUND")
		}
	} else {
		fmt.Println("GET USER BY ID " + id)
		if data := service.GetUserById(id, user); data != nil {
			ctx.JSON(data)
		} else {
			ctx.StatusCode(404)
		}
	}
}

func AddUser(ctx iris.Context, service Connection, user *User) {
	rawJson, err := ctx.GetBody()
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	if data := service.AddUser(rawJson, user); data != nil {
		ctx.JSON(data)
	} else {
		ctx.StatusCode(500)
	}
}

func UpdateUser(ctx iris.Context, service Connection, user *User) {
	if rawJson, err := ctx.GetBody(); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	} else {
		if data := service.UpdateUser(rawJson, user); data != nil {
			ctx.JSON(data)
		} else {
			ctx.StatusCode(500)
		}
	}
}

func DeleteUser(ctx iris.Context, service Connection, user *User) {
	id := ctx.Params().Get("id")
	if num := service.DeleteUser(id, user); num == 1 {
		ctx.JSON(iris.Map{"id": id})
	} else {
		ctx.StatusCode(404)
	}
}
