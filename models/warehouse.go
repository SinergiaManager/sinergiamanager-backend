package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemWarehouse struct {
	ItemDb   primitive.ObjectID `json:"item_id" bson:"item_id" validate:"required,mongodb"`
	Quantity int                `json:"quantity" bson:"quantity" validate:"required,gte=0"`
}

type WarehouseDb struct {
	ID       string          `bson:"_id"`
	Name     string          `bson:"name"`
	Location string          `bson:"location"`
	Code     string          `bson:"code, omitempty"`
	Items    []ItemWarehouse `bson:"items, omitempty"`
	UpdateAt time.Time       `bson:"update_at"`
	InsertAt time.Time       `bson:"insert_at"`
}

type WarehouseIns struct {
	Name     string          `json:"name" bson:"name" validate:"required,gte=3,lte=50"`
	Location string          `json:"location" bson:"location" validate:"required,gte=2,lte=50"`
	Code     string          `json:"code" bson:"code,omitempty"`
	Items    []ItemWarehouse `json:"items" bson:"items,omitempty" validate:"dive"`
	UpdateAt time.Time       `json:"update_at" bson:"update_at"`
	InsertAt time.Time       `json:"insert_at" bson:"insert_at"`
}
