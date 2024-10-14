package models

import "time"

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

type WarehouseIn struct {
	Name     string          `json:"name" validate:"required"`
	Location string          `json:"location" validate:"required"`
	Code     string          `json:"code" validate:"required"`
	Items    []ItemWarehouse `json:"items"`
	UpdateAt time.Time       `json:"update_at"`
	InsertAt time.Time       `json:"insert_at"`
}
