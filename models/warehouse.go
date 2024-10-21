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
	Name     string          `json:"name" bson:"name" validate:"required"`
	Location string          `json:"location" bson:"location" validate:"required"`
	Code     string          `json:"code" bson:"code"`
	Items    []ItemWarehouse `json:"items" bson:"items"`
	UpdateAt time.Time       `json:"update_at" bson:"update_at"`
	InsertAt time.Time       `json:"insert_at" bson:"insert_at"`
}

func WarehouseStructLevelValidation(sl validator.StructLevel) {
	warehouse := sl.Current().Interface().(WarehouseIns)

	/* input validation */
	if len(warehouse.Name) < 3 {
		sl.ReportError(warehouse.Name, "Name", "name", "minlength", "3")
	}
}
