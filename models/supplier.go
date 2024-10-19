package models

import (
	"time"
)

type SupplierIns struct {
	Name     string    `json:"name" bson:"name" validate:"required"`
	Address  string    `json:"address" bson:"address" validate:"required"`
	Phone    string    `json:"phone" bson:"phone" validate:"required,e164"`
	Code     string    `json:"code" bson:"code"`
	InsertAt time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
}

type SupplierDb struct {
	ID             string    `bson:"_id"`
	Name           string    `bson:"name"`
	Address        string    `bson:"address"`
	Phone          string    `bson:"phone"`
	Code           string    `bson:"code"`
	ItemSuppliable []string  `bson:"item_suppliable"`
	InsertAt       time.Time `bson:"insert_at"`
	UpdateAt       time.Time `bson:"update_at"`
}
