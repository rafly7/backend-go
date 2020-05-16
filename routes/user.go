package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	ctr "github.com/rafly7/backend/controllers"
	"github.com/rafly7/backend/middlewares"
	. "github.com/rafly7/backend/models"
	. "github.com/rafly7/backend/services"
)

func UserRoutes(router iris.Party, db *gorm.DB) {
	//ss := services.Connection{Db: db}
	// users := &models.Users{}
	// fmt.Println(users)
	// ctx.JSON(ss.GetAllUser(users))
	router.Use(middlewares.JwtHandler().Serve)
	router.Get("/", func(ctx iris.Context) { ctr.GetUser(ctx, Connection{Db: db}, &User{}, &Users{}) })
	router.Post("/", func(ctx iris.Context) { ctr.AddUser(ctx, Connection{Db: db}, &User{}) })
	router.Put("/", func(ctx iris.Context) { ctr.UpdateUser(ctx, Connection{Db: db}, &User{}) })
	router.Delete("/{id:string}", func(ctx iris.Context) { ctr.DeleteUser(ctx, Connection{Db: db}, &User{}) })
	// router.Get("/{id:string}", func(ctx iris.Context) {
	// 	id := ctx.Params().Get("id")
	// 	exists := ss.GetUserById(id, &models.User{})
	// 	if exists != nil {
	// 		ctx.StatusCode(200)
	// 		ctx.JSON(exists)
	// 	} else {
	// 		ctx.StatusCode(404)
	// 		ctx.JSON(iris.Map{"message": "Not Found"})
	// 	}
	// })
	// router.Post("/", func(ctx iris.Context) {
	// 	rawJson, err := ctx.GetBody()
	// 	if err != nil { /* handle the error */
	// 		ctx.Writef("%v", err)
	// 	}
	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(ss.AddUser(rawJson, &models.User{}))
	// })
	// router.Put("/", func(ctx iris.Context) {
	// 	body := ctx.Request().Body
	// 	rawBodyAsBytes, err := ioutil.ReadAll(body)
	// 	if err != nil {
	// 		ctx.Writef("%v", err)
	// 	}
	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(ss.UpdateUser(rawBodyAsBytes, &models.User{}))
	// 	defer body.Close()
	// 	// ctx.JSON(iris.Map{"message": "Put request method"})
	// })
	// router.Delete("/{id:string}", func(ctx iris.Context) {
	// 	id := ctx.Params().Get("id")
	// 	ctx.JSON(ss.DeleteUser(id, &models.User{}))
	// })
}
