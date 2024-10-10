package main

import (
	"log"

	ConfigAuth "github.com/SinergiaManager/sinergiamanager-backend/config/auth"
	ConfigDb "github.com/SinergiaManager/sinergiamanager-backend/config/database"
	Controllers "github.com/SinergiaManager/sinergiamanager-backend/controllers"

	"github.com/kataras/iris/v12"
)

func main() {
	if err := ConfigDb.ConnectDb(); err != nil {
		ConfigDb.DisconnectDb()
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer ConfigDb.DisconnectDb()

	app := iris.New()

	user := app.Party("/user")
	{
		user.Get("/", Controllers.GetAllUsers)
		user.Post("/", Controllers.CreateUser)
		user.Delete("/{id:string}", Controllers.DeleteUser)
	}

	user.Use(ConfigAuth.Protected)

	item := app.Party("/item")
	{
		item.Get("/", Controllers.GetAllItems)
		item.Post("/", Controllers.CreateItem)
		item.Delete("/{id:string}", Controllers.DeleteItem)
	}

	auth := app.Party("/auth")
	{
		auth.Post("/login", Controllers.Login)
		auth.Post("/logout", Controllers.Logout)
		//auth.Post("/register", Controllers.Register)
	}

	app.Listen(":8080")
}
