package services

import (
	"context"
	"math/rand"
	"time"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Seeder() {
	Config.DB.Drop(context.Background())

	clients := []Models.ClientIns{}
	warehouses := []Models.WarehouseIns{}
	ctx := context.TODO()

	for i := 0; i < 5; i++ {
		client := Models.ClientIns{
			Name:     faker.Name(),
			Surname:  faker.LastName(),
			Email:    faker.Email(),
			Phone:    faker.Phonenumber(),
			Address:  faker.GetRealAddress().Address,
			InsertAt: time.Now().UTC(),
			UpdateAt: time.Now().UTC(),
		}
		clients = append(clients, client)
	}

	clientInterfaces := make([]interface{}, len(clients))
	for i, v := range clients {
		clientInterfaces[i] = v
	}
	_, err := Config.DB.Collection("clients").InsertMany(ctx, clientInterfaces)
	if err != nil {
		panic(err)
	}

	supplierIDs := []primitive.ObjectID{}
	for i := 0; i < 5; i++ {
		supplier := Models.SupplierIns{
			Name:      faker.Name(),
			Email:     faker.Email(),
			Phone:     faker.Phonenumber(),
			Address:   faker.GetRealAddress().Address,
			VATNumber: "DE100000000",
			InsertAt:  time.Now().UTC(),
			UpdateAt:  time.Now().UTC(),
		}
		result, err := Config.DB.Collection("suppliers").InsertOne(ctx, supplier)
		if err != nil {
			panic(err)
		}
		supplierIDs = append(supplierIDs, result.InsertedID.(primitive.ObjectID))
	}

	ItemIds := []primitive.ObjectID{}
	for i := 0; i < 10; i++ {
		item := Models.ItemIns{
			Name:        faker.Word(),
			Code:        faker.UUIDHyphenated(),
			SupplierID:  supplierIDs[rand.Intn(len(supplierIDs))], // Random supplier ObjectID
			Description: faker.Sentence(),
			InsertAt:    time.Now().UTC(),
			UpdateAt:    time.Now().UTC(),
		}
		result, err := Config.DB.Collection("items").InsertOne(ctx, item)
		if err != nil {
			panic(err)
		}
		ItemIds = append(ItemIds, result.InsertedID.(primitive.ObjectID))
	}

	for i := 0; i < 3; i++ {
		warehouseItems := []Models.ItemWarehouse{}
		for j := 0; j < 3; j++ {
			warehouseItems = append(warehouseItems, Models.ItemWarehouse{
				ItemDb:   ItemIds[rand.Intn(len(ItemIds))], // Random item ObjectID
				Quantity: rand.Intn(100) + 1,               // Random quantity
			})
		}

		warehouse := Models.WarehouseIns{
			Name:     faker.Word(),
			Location: faker.GetRealAddress().Address,
			Code:     faker.UUIDHyphenated(),
			Items:    warehouseItems,
			InsertAt: time.Now().UTC(),
			UpdateAt: time.Now().UTC(),
		}
		warehouses = append(warehouses, warehouse)
	}

	warehouseInterfaces := make([]interface{}, len(warehouses))
	for i, v := range warehouses {
		warehouseInterfaces[i] = v
	}
	_, err = Config.DB.Collection("warehouses").InsertMany(ctx, warehouseInterfaces)
	if err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	userAdmin := bson.M{
		"_id": func() primitive.ObjectID {
			id, err := primitive.ObjectIDFromHex("673794884216b36623876602")
			if err != nil {
				panic(err)
			}
			return id
		}(),
		"username":  "admin",
		"name":      "Admin",
		"surname":   "Admin",
		"email":     "admin@admin.com",
		"password":  string(hashedPassword),
		"role":      "admin",
		"insert_at": time.Now().UTC(),
		"update_at": time.Now().UTC(),
	}

	result, err := Config.DB.Collection("users").InsertOne(ctx, userAdmin)
	if err != nil {
		panic(err)
	}

	adminID := result.InsertedID.(primitive.ObjectID)

	for i := 0; i < 5; i++ {
		punchRecord := Models.PunchRecordIns{
			EmployeeID: adminID,
			PunchType: func() string {
				if rand.Intn(2) == 0 {
					return "IN"
				} else {
					return "OUT"
				}
			}(),
			PunchTime: time.Now(),
			Location:  faker.GetRealAddress().Address,
			InsertAt:  time.Now().UTC(),
			UpdateAt:  time.Now().UTC(),
		}
		_, err := Config.DB.Collection("punch_records").InsertOne(ctx, punchRecord)
		if err != nil {
			panic(err)
		}
	}
}
