package main

import (
	"log"

	ConfigAuth "github.com/SinergiaManager/sinergiamanager-backend/config/auth"
	ConfigDb "github.com/SinergiaManager/sinergiamanager-backend/config/database"
	Enum "github.com/SinergiaManager/sinergiamanager-backend/config/utils"
	Controllers "github.com/SinergiaManager/sinergiamanager-backend/controllers"
	UserModel "github.com/SinergiaManager/sinergiamanager-backend/models"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func main() {
	if err := ConfigDb.ConnectDb(); err != nil {
		ConfigDb.DisconnectDb()
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer ConfigDb.DisconnectDb()

	ConfigAuth.InitJWT()

	v := validator.New()
	v.RegisterStructValidation(UserModel.UserStructLevelValidation, UserModel.UserIns{})

	app := iris.New()
	app.Validator = v

	user := app.Party("/users")
	{
		user.Get("/", ConfigAuth.JWTMiddleware([]string{string(Enum.EnumUserRole.ADMIN)}), Controllers.GetAllUsers)
		user.Post("/", ConfigAuth.JWTMiddleware([]string{string(Enum.EnumUserRole.ADMIN)}), Controllers.CreateUser)
		user.Put("/{id:string}", ConfigAuth.JWTMiddleware([]string{}), Controllers.UpdateUser)
		user.Delete("/{id:string}", ConfigAuth.JWTMiddleware([]string{}), Controllers.DeleteUser)
	}

	item := app.Party("/items")
	{
		item.Get("/", Controllers.GetAllItems)
		item.Post("/", Controllers.CreateItem)
		item.Put("/{id:string}", Controllers.UpdateItem)
		item.Delete("/{id:string}", Controllers.DeleteItem)
	}

	auth := app.Party("/auth")
	{
		auth.Post("/login", Controllers.Login)
		auth.Post("/logout", ConfigAuth.JWTMiddleware([]string{}), Controllers.Logout)
		auth.Post("/register", Controllers.Register)
	}

	app.Listen(":8080")
}
