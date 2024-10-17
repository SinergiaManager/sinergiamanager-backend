package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ItemWarehouse struct {
	ItemDb   string `bson:"item_id"`
	Quantity int    `bson:"quantity"`
}

type WarehouseDb struct {
	ID       string          `bson:"_id"`
	Name     string          `bson:"name"`
	Location string          `bson:"location"`
	Code     string          `bson:"code"`
	Items    []ItemWarehouse `bson:"items"`
	UpdateAt time.Time       `bson:"update_at"`
	InsertAt time.Time       `bson:"insert_at"`
}

type WarehouseIns struct {
	Name     string          `json:"name" validate:"required"`
	Location string          `json:"location" validate:"required"`
	Code     string          `json:"code" `
	Items    []ItemWarehouse `json:"items"`
	UpdateAt time.Time       `json:"update_at"`
	InsertAt time.Time       `json:"insert_at"`
}

func WarehouseStructLevelValidation(wl validator.StructLevel) {
	warehouse := wl.Current().Interface().(WarehouseIns)

	if len(warehouse.Name) > 3 {
		wl.ReportError(warehouse.Name, "Name", "Name", "name", "")
	}
}
