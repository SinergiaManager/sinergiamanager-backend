package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ItemDb struct {
	ID          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Code        string    `bson:"code"`
	SupplierID  string    `bson:"supplier_id"`
	Description string    `bson:"description"`
	InsertAt    time.Time `bson:"insert_at"`
	UpdateAt    time.Time `bson:"update_at"`
}

type ItemIns struct {
	Name        string    `json:"name" bson:"name" validate:"required,gte=3,lte=50"`
	Code        string    `json:"code" bson:"code"`
	SupplierID  string    `json:"supplier_id" bson:"supplier_id"`
	Description string    `json:"description" bson:"description" validate:"gte=3,lte=100"`
	InsertAt    time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt    time.Time `json:"update_at" bson:"update_at"`
}

func ItemStructLevelValidation(sl validator.StructLevel) {
	//item := sl.Current().Interface().(ItemIns)
}
