package main

import (
	"log"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Controllers "github.com/SinergiaManager/sinergiamanager-backend/controllers"

	"github.com/kataras/iris/v12"
)

func main() {
	if err := Config.ConnectDb(); err != nil {
		Config.DisconnectDb()
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer Config.DisconnectDb()

	app := iris.New()

	user := app.Party("/user")
	{
		user.Get("/", Controllers.GetAllUsers)
		user.Post("/", Controllers.CreateUser)
		user.Delete("/{id:string}", Controllers.DeleteUser)
	}

	item := app.Party("/item")
	{
		item.Get("/", Controllers.GetAllItems)
	}

	app.Listen(":8080")
}
