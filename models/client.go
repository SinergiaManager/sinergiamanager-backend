package models

import (
	"time"
)

type ClientIns struct {
	Name     string    `json:"name" bson:"name" validate:"required"`
	Email    string    `json:"email" bson:"email" validate:"required,email"`
	Phone    string    `json:"phone" bson:"phone" validate:"required,e164"`
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
