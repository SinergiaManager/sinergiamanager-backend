package services

import (
	"time"

	"context"

	"strconv"

	Config "github.com/SinergiaManager/sinergiamanager-backend/config"
	Models "github.com/SinergiaManager/sinergiamanager-backend/models"
	"github.com/go-faker/faker/v4"
)

func Seeder() {
	clients := []Models.ClientIns{}
	items := []Models.ItemIns{}
	punchRecords := []Models.PunchRecordIns{}
	suppliers := []Models.SupplierIns{}
	warehouses := []Models.WarehouseIns{}

	ctx := context.TODO()
	for i := 0; i < 10; i++ {
		client := Models.ClientIns{
			Name:     faker.Name(),
			Surname:  faker.LastName(),
			Email:    faker.Email(),
			Phone:    faker.Phonenumber(),
			Address:  faker.GetRealAddress().Address,
			InsertAt: time.Now(),
			UpdateAt: time.Now(),
		}

		clients = append(clients, client)

		vat, err := faker.RandomInt(329494338, 999999999)
		if err != nil {
			panic(err)
		}

		supplier := Models.SupplierIns{
			Name:      faker.Name(),
			Email:     faker.Email(),
			Phone:     faker.Phonenumber(),
			Address:   faker.GetRealAddress().Address,
			VATNumber: "DE" + strconv.Itoa(vat[0]),
			InsertAt:  time.Now(),
			UpdateAt:  time.Now(),
		}

		suppliers = append(suppliers, supplier)

	}

	clientInterfaces := make([]interface{}, len(clients))
	for i, v := range clients {
		clientInterfaces[i] = v
	}

	_, err := Config.DB.Collection("clients").InsertMany(ctx, clientInterfaces)
	if err != nil {
		panic(err)
	}

	item := Models.ItemIns{
		Name:        faker.Name(),
		Code:        faker.UUIDDigit(),
		SupplierID:  faker.UUIDDigit(),
		Description: faker.Sentence(),
		InsertAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
}
