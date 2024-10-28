package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ltns35/go-vat"
)

type SupplierIns struct {
	Name      string    `json:"name" bson:"name" validate:"required, gte=3"`
	Address   string    `json:"address" bson:"address" validate:"required"`
	Phone     string    `json:"phone" bson:"phone" validate:"e164"`
	Email     string    `json:"email" bson:"email" validate:"email"`
	VATNumber string    `json:"vat_number" bson:"vat_number"`
	Code      string    `json:"code" bson:"code"`
	InsertAt  time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt  time.Time `json:"update_at" bson:"update_at"`
}

type SupplierDb struct {
	ID             string    `bson:"_id"`
	Name           string    `bson:"name"`
	Address        string    `bson:"address"`
	Phone          string    `bson:"phone"`
	Email          string    `bson:"email"`
	Code           string    `bson:"code"`
	VATNumber      string    `bson:"vat_number"`
	ItemSuppliable []string  `bson:"item_suppliable"`
	InsertAt       time.Time `bson:"insert_at"`
	UpdateAt       time.Time `bson:"update_at"`
}

func SupplierstructLevelValidation(wl validator.StructLevel) {
	supplier := wl.Current().Interface().(SupplierIns)

	vatResult, err := vat.Validate(supplier.VATNumber)
	if err != nil {
		wl.ReportError(supplier.VATNumber, "VATNumber", "VATNumber", "VATNumber", "")
	}

	if !vatResult.IsValid {
		wl.ReportError(supplier.VATNumber, "VATNumber", "VATNumber", "VATNumber", "")
	}
}
