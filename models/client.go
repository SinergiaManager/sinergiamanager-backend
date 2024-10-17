package models

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type ClientIns struct {
	Name     string    `json:"name" bson:"name" validate:"required"`
	Email    string    `json:"email" bson:"email" validate:"required"`
	Phone    string    `json:"phone" bson:"phone" validate:"required"`
	Address  string    `json:"address" bson:"address" validate:"required"`
	InsertAt time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
}

type ClientDb struct {
	ID       string    `bson:"_id"`
	Name     string    `bson:"name"`
	Email    string    `bson:"email"`
	Phone    string    `bson:"phone"`
	Address  string    `bson:"address"`
	InsertAt time.Time `bson:"insert_at"`
	UpdateAt time.Time `bson:"update_at"`
}

func ClientStructLevelValidation(sl validator.StructLevel) {
	client := sl.Current().Interface().(ClientIns)

	match := `^[\w\-\.\d]+@[\w\-\.\d]+\.[\w\-\.\d]+$`
	if ok, _ := regexp.MatchString(match, client.Email); !ok {
		sl.ReportError(client.Email, "Email", "Email", "email", "")
	}

	match = `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`
	if ok, _ := regexp.MatchString(match, client.Phone); !ok {
		sl.ReportError(client.Phone, "Phone", "Phone", "phone", "")
	}
}
