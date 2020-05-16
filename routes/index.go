package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

func IndexRoutes(app *iris.Application, db *gorm.DB) {
	AuthRoutes(app.Party("/auth"), db)
	UserRoutes(app.Party("/user"), db)
	PermissionRoutes(app.Party("/permission"), db)
}
