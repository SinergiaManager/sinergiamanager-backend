package main

import (
	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Controllers "github.com/SinergiaManager/sinergiamanager-backend/controllers"

	"github.com/kataras/iris/v12"
)

func main() {
	Config.ConnectDb()
	defer Config.DisconnectDb()
	app := iris.New()

	user := app.Party("/user")
	{
		user.Get("/", Controllers.GetAllUsers)
	}

	app.Listen(":8080")
}
