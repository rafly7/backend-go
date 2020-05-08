package servers

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func Server(level string, db *gorm.DB) *iris.Application {
	//f := events.LogFile()
	//defer f.Close()

	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel(level)
	//app.Logger().SetOutput(f)
	app.Use(requestLogger)
	//mvc.New(routes.IndexRoutes(app, db))
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("ping")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"msg": "Hello Iris!"})
	})
	//routes.IndexRoutes(app, db)
	return app
}
