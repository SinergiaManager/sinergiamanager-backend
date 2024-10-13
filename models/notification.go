package models

import (
	"time"
)

type Notification struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserID      string    `json:"userId" bson:"userId"` // user id of the receiver
	Title       string    `json:"title" bson:"title"`
	Message     string    `json:"message" bson:"message"`
	Types       []string  `json:"types" bson:"types"`
	IsDelivered bool      `json:"isDelivered" bson:"isDelivered"`
	InsertAt    time.Time `bson:"insert_at"`
	UpdateAt    time.Time `bson:"update_at"`
}
