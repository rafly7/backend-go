package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	ctr "github.com/rafly7/backend/controllers"
	"github.com/rafly7/backend/models"
	. "github.com/rafly7/backend/services"
)

func AuthRoutes(router iris.Party, db *gorm.DB) {
	router.Post("/", func(ctx iris.Context) { ctr.GetToken(ctx, Connection{Db: db}, &models.User{}) })
}
