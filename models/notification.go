package models

import (
	"time"
)

type NotificationDb struct {
	ID          string    `bson:"_id,omitempty"`
	UserID      string    `bson:"user_id"` // user id of the receiver
	Title       string    `bson:"title"`
	Message     string    `bson:"message"`
	Types       []string  `bson:"types"`
	IsDelivered bool      `bson:"is_delivered"`
	IsRead      bool      `bson:"is_read"`
	InsertAt    time.Time `bson:"insert_at"`
	UpdateAt    time.Time `bson:"update_at"`
	DeliveredAt time.Time `bson:"delivered_at"`
	ReadAt      time.Time `bson:"read_at"`
}

type NotificationIns struct {
	UserID      string    `json:"user_id" bson:"user_id" validate:"required"` // user id of the receiver
	Title       string    `json:"title" bson:"title" validate:"required"`
	Message     string    `json:"message" bson:"message" validate:"required"`
	Types       []string  `json:"types" bson:"types" validate:"required"`
	IsDelivered bool      `json:"is_delivered" bson:"is_delivered" validate:"required"`
	IsRead      bool      `json:"is_read" bson:"is_read" validate:"required"`
	DeliveredAt time.Time `json:"delivered_at" bson:"delivered_at" validate:"required"`
	ReadAt      time.Time `json:"read_at" bson:"read_at" validate:"required"`
	InsertAt    time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt    time.Time `json:"update_at" bson:"update_at"`
}
