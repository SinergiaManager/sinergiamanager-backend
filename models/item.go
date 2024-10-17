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
	Name        string    `json:"name" bson:"name" validate:"required"`
	Code        string    `json:"code" bson:"code"`
	SupplierID  string    `json:"supplier_id" bson:"supplier_id"`
	Description string    `json:"description" bson:"description"`
	InsertAt    time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt    time.Time `json:"update_at" bson:"update_at"`
}

func ItemStructLevelValidation(sl validator.StructLevel) {
	item := sl.Current().Interface().(ItemIns)

	/* input validation */
	if len(item.Name) <= 0 {
		sl.ReportError(item.Name, "Name", "name", "name", "")
	}
}
