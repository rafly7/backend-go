package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/rafly7/backend/configs"
	"github.com/rafly7/backend/servers"
)

func main() {

	// Command for backend
	// migrateCommand := flag.NewFlagSet("migrate", flag.ExitOnError)
	// runCommand := flag.NewFlagSet("running", flag.ExitOnError)
	// if len(os.Args) < 2 {
	// 	println("migrate or running is required")
	// 	os.Exit(1)
	// }
	// switch os.Args[1] {
	// case "migrate":
	// 	migrateCommand.Parse(os.Args[2:])
	// case "running":
	// 	runCommand.Parse(os.Args[2:])
	// default:
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	// // Command for migrations database
	// if migrateCommand.Parsed() {
	// 	db, err := configs.ConnectToDb()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		panic("failed to connect database")
	// 	}
	// 	defer db.Close()
	// 	db.SingularTable(true)
	// 	db.LogMode(true)
	// 	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// 	migrations.RunMigrations(db)
	// }

	// // Command for running server
	// if runCommand.Parsed() {
	db, err := configs.ConnectToDb()
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	address, level := configs.ConfigServer()
	app, f := servers.Server(level, db)
	defer f.Close()
	app.Run(
		iris.Addr(address),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
	//}
}
