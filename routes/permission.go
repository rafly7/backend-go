package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	ctr "github.com/rafly7/backend/controllers"
	. "github.com/rafly7/backend/models"
	. "github.com/rafly7/backend/services"
)

func PermissionRoutes(router iris.Party, db *gorm.DB) {
	// router.Use(middlewares.JwtHandler().Serve)
	router.Get("/", func(ctx iris.Context) { ctr.GetPermission(ctx, Connection{Db: db}, &Permission{}, &Permissions{}) })
	router.Post("/", func(ctx iris.Context) { ctr.AddPermission(ctx, Connection{Db: db}, &Permission{}) })
	router.Put("/", func(ctx iris.Context) { ctr.UpdatePermission(ctx, Connection{Db: db}, &Permission{}) })
	router.Delete("/{id:string}", func(ctx iris.Context) { ctr.DeletePermission(ctx, Connection{Db: db}, &Permission{}) })
}
