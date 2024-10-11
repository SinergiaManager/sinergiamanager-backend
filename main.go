package main

import (
	"log"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Controllers "github.com/SinergiaManager/sinergiamanager-backend/controllers"
	UserModel "github.com/SinergiaManager/sinergiamanager-backend/models"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func main() {
	if err := Config.ConnectDb(); err != nil {
		Config.DisconnectDb()
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer Config.DisconnectDb()

	Config.InitJWT()

	v := validator.New()
	v.RegisterStructValidation(UserModel.UserStructLevelValidation, UserModel.UserIns{})

	app := iris.New()
	app.Validator = v

	user := app.Party("/users")
	{
		user.Get("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetAllUsers)
		user.Get("/me", Config.JWTMiddleware([]string{}), Controllers.GetMe)
		user.Get("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetUser)
		user.Post("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.CreateUser)
		user.Put("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.UpdateUser)
		user.Put("/me", Config.JWTMiddleware([]string{}), Controllers.UpdateMe)
		user.Delete("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.DeleteUser)
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
		auth.Post("/logout", Config.JWTMiddleware([]string{}), Controllers.Logout)
		auth.Post("/register", Controllers.Register)
		auth.Post("/refresh", Config.JWTMiddleware([]string{}), Controllers.RefreshToken)
	}

	app.Listen(":8080")
}
