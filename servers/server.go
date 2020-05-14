package servers

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/rafly7/backend/events"
	"github.com/rafly7/backend/routes"
)

func Server(level string, db *gorm.DB) (*iris.Application, *os.File) {
	f := events.LogFile()
	// defer f.Close()

	app := iris.New()
	// app.Use(recover.New())
	app.Logger().SetLevel(level)
	//app.Logger().SetOutput(f)
	app.Use(requestLogger)
	routes.IndexRoutes(app, db)
	// mvc.Configure(app.Party("/user"), routes.user)
	//mvc.New(routes.IndexRoutes(app, db))
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome guys</h1>")
	})
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("ping")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"msg": "Hello Iris!"})
	})
	//routes.IndexRoutes(app, db)
	return app, f
}
