package main

import (
	"context"
	"log"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Controllers "github.com/SinergiaManager/sinergiamanager-backend/controllers"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	Services "github.com/SinergiaManager/sinergiamanager-backend/services"

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
	v.RegisterStructValidation(Models.UserChangePasswordStructLevelValidation, Models.UserChangePassword{})
	v.RegisterStructValidation(Models.ItemStructLevelValidation, Models.ItemIns{})

	app := iris.New()
	app.Validator = v

	user := app.Party("/users")
	{
		user.Get("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetAllUsers)
		user.Get("/me", Config.JWTMiddleware([]string{}), Controllers.GetMe)
		user.Get("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetUser)

		user.Post("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.CreateUser)
		user.Post("/change-password", Config.JWTMiddleware([]string{}), Controllers.ChangePassword)
		user.Post("/forgot-password", Config.JWTMiddleware([]string{}), Controllers.ForgotPassword)

		user.Put("/me", Config.JWTMiddleware([]string{}), Controllers.UpdateMe)
		user.Put("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.UpdateUser)

		user.Delete("/me", Config.JWTMiddleware([]string{}), Controllers.DeleteMe)
		user.Delete("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.DeleteUser)
	}

	item := app.Party("/items")
	{
		item.Get("/", Controllers.GetAllItems)
		item.Post("/", Controllers.CreateItem)
		item.Put("/{id:string}", Controllers.UpdateItem)
		item.Delete("/{id:string}", Controllers.DeleteItem)
	}

	warehouse := app.Party("/warehouses")
	{
		warehouse.Get("/", Controllers.GetAllWarehouses)
		warehouse.Get("/{id:string}", Controllers.GetWarehouseById)
		warehouse.Post("/", Controllers.CreateWarehouse)
		warehouse.Put("/{id:string}", Controllers.UpdateWarehouse)
		warehouse.Delete("/{id:string}", Controllers.DeleteWarehouse)
	}

	auth := app.Party("/auth")
	{
		auth.Post("/login", Controllers.Login)
		auth.Post("/logout", Config.JWTMiddleware([]string{}), Controllers.Logout)
		auth.Post("/register", Controllers.Register)
		auth.Post("/refresh", Config.JWTMiddleware([]string{}), Controllers.RefreshToken)
	}

	notification := app.Party("/notifications")
	{
		notification.Get("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetAllNotifications)
		notification.Get("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetNotification)
		notification.Get("/user/{userID:string}", Config.JWTMiddleware([]string{}), Controllers.GetNotificationsByUser)
		notification.Get("/me", Config.JWTMiddleware([]string{}), Controllers.GetNotificationsMe)
		notification.Get("/me/{id:string}", Config.JWTMiddleware([]string{}), Controllers.GetNotificationMe)

		notification.Post("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.CreateNotification)

		notification.Put("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.UpdateNotification)

		notification.Delete("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.DeleteNotification)
	}

	supplier := app.Party("/suppliers")
	{
		supplier.Get("/", Config.JWTMiddleware([]string{}), Controllers.GetAllSuppliers)
		supplier.Get("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.GetSupplier)
		supplier.Post("/", Config.JWTMiddleware([]string{}), Controllers.CreateSupplier)
		supplier.Put("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.UpdateSupplier)
		supplier.Delete("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.DeleteSupplier)
	}

	client := app.Party("/clients")
	{
		client.Get("/", Config.JWTMiddleware([]string{}), Controllers.GetAllClients)
		client.Get("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.GetClient)
		client.Post("/", Config.JWTMiddleware([]string{}), Controllers.CreateClient)
		client.Put("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.UpdateClient)
		client.Delete("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.DeleteClient)
	}

	go Services.SetupJobScheduler(context.TODO())

	app.Listen(":8080")
}
