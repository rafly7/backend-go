package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

func IndexRoutes(app *iris.Application, db *gorm.DB) {
	UserRoutes(app.Party("/user"), db)
}
