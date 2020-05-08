package routes

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/rafly7/backend/models"
	"github.com/rafly7/backend/services"
)

func UserRoutes(router iris.Party, db *gorm.DB) {
	ss := services.Connection{Db: db}
	users := &models.Users{}
	router.Get("/", func(ctx iris.Context) {
		fmt.Println(users)
		ctx.JSON(ss.GetAllUser(users))
	})
	router.Get("/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		exists := ss.GetUserById(id, &models.User{})
		if exists != nil {
			ctx.StatusCode(200)
			ctx.JSON(exists)
		} else {
			ctx.StatusCode(404)
			ctx.JSON(iris.Map{"message": "Not Found"})
		}
	})
	router.Post("/", func(ctx iris.Context) {
		body := ctx.Request().Body
		rawBodyAsBytes, err := ioutil.ReadAll(body)
		if err != nil { /* handle the error */
			ctx.Writef("%v", err)
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(ss.AddUser(rawBodyAsBytes, &models.User{}))
		defer body.Close()
	})
	router.Put("/", func(ctx iris.Context) {
		body := ctx.Request().Body
		rawBodyAsBytes, err := ioutil.ReadAll(body)
		if err != nil {
			ctx.Writef("%v", err)
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(ss.UpdateUser(rawBodyAsBytes, &models.User{}))
		defer body.Close()
		// ctx.JSON(iris.Map{"message": "Put request method"})
	})
	router.Delete("/{id:string}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		ctx.JSON(ss.DeleteUser(id, &models.User{}))
	})
}
