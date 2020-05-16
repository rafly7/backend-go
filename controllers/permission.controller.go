package controllers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	. "github.com/rafly7/backend/models"
	. "github.com/rafly7/backend/services"
)

func GetPermission(ctx iris.Context, service Connection, permission *Permission, permissions *Permissions) {
	id := ctx.URLParam("id")
	if len(id) == 0 {
		fmt.Println("GET ALL PERMISSION")
		if data := service.GetAllPermission(permissions); data != nil {
			ctx.JSON(data)
		} else {
			ctx.StatusCode(404)
			iris.New().Logger().Errorf("%s", "ERROR NOT FOUND")
		}
	} else {
		fmt.Println("GET PERMISSION BY ID " + id)
		if data := service.GetPermissionById(id, permission); data != nil {
			ctx.JSON(data)
		} else {
			ctx.StatusCode(404)
		}
	}
}

func AddPermission(ctx iris.Context, service Connection, permission *Permission) {
	rawJson, err := ctx.GetBody()
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	if data := service.AddPermission(rawJson, permission); data != nil {
		ctx.JSON(data)
	} else {
		ctx.StatusCode(500)
	}
}

func UpdatePermission(ctx iris.Context, service Connection, permission *Permission) {
	if rawJson, err := ctx.GetBody(); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	} else {
		if data := service.UpdatePermission(rawJson, permission); data != nil {
			ctx.JSON(data)
		} else {
			ctx.StatusCode(500)
		}
	}
}

func DeletePermission(ctx iris.Context, service Connection, permission *Permission) {
	id := ctx.Params().Get("id")
	if num := service.DeletePermission(id, permission); num == 1 {
		ctx.JSON(iris.Map{"id": id})
	} else {
		ctx.StatusCode(404)
	}
}
