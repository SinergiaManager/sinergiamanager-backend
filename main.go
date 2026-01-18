package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SinergiaManager/sinergiamanager-backend/internal/warehouse"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kataras/iris/v12"
)

func main() {
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	defer db.Close()
	/* if err := Config.ConnectDb(); err != nil {
		Config.DisconnectDb()
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	Config.DockerConfig()

	defer Config.DisconnectDb()

	Services.Seeder()

	Config.InitJWT()

	v := validator.New()
	v.RegisterStructValidation(Models.UserChangePasswordStructLevelValidation, Models.UserChangePassword{})
	v.RegisterStructValidation(Models.ItemStructLevelValidation, Models.ItemIns{})
	v.RegisterStructValidation(Models.PunchRecordStructLevelValidation, Models.PunchRecordIns{})
	v.RegisterStructValidation(Models.SupplierStructLevelValidation, Models.SupplierIns{})
	v.Struct(Models.WarehouseIns{}) */

	app := iris.New()

	routes := app.Party("/")
	warehouseModule := warehouse.NewModule(db)
	warehouseModule.RegisterRoutes(routes)

	/* app.Validator = v */

	/* config := app.Party("/configs")
	{
		config.Get("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetAllConfigs)
		config.Get("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.GetConfig)
		config.Post("/", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.CreateConfig)
		config.Put("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.UpdateConfig)
		config.Delete("/{id:string}", Config.JWTMiddleware([]string{string(Config.EnumUserRole.ADMIN)}), Controllers.DeleteConfig)
	}

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
	} */

	/* warehouse := app.Party("/warehouses")
	{
		warehouse.Get("/", Controllers.GetAllWarehouses)
		warehouse.Get("/{id:string}", Controllers.GetWarehouseById)
		warehouse.Post("/", Controllers.CreateWarehouse)
		warehouse.Put("/{id:string}", Controllers.UpdateWarehouse)
		warehouse.Delete("/{id:string}", Controllers.DeleteWarehouse)
	} */

	/* auth := app.Party("/auth")
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
		notification.Get("/sse", Config.JWTMiddleware([]string{}), Controllers.GetNotificationSSEMe)

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

	punchRecord := app.Party("/punch-records")
	{
		punchRecord.Get("/", Config.JWTMiddleware([]string{}), Controllers.GetAllPunchRecords)
		punchRecord.Get("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.GetPunchRecord)
		punchRecord.Post("/", Config.JWTMiddleware([]string{}), Controllers.CreatePunchRecord)
		punchRecord.Put("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.UpdatePunchRecord)
		punchRecord.Delete("/{id:string}", Config.JWTMiddleware([]string{}), Controllers.DeletePunchRecord)
	} */

	/* go Services.SetupJobScheduler(context.TODO()) */

	/* crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Authorization"},
	})
	app.UseRouter(crs) */

	app.Listen(":8080")
}
