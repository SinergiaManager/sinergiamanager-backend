package models

import (
	"time"
)

type NotificationDb struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserID      string    `json:"userId" bson:"userId"` // user id of the receiver
	Title       string    `json:"title" bson:"title"`
	Message     string    `json:"message" bson:"message"`
	Types       []string  `json:"types" bson:"types"`
	IsDelivered bool      `json:"isDelivered" bson:"isDelivered"`
	IsRead      bool      `json:"isRead" bson:"isRead"`
	InsertAt    time.Time `bson:"insert_at"`
	UpdateAt    time.Time `bson:"update_at"`
	DeliveredAt time.Time `json:"deliveredAt" bson:"deliveredAt"`
	ReadAt      time.Time `json:"readAt" bson:"readAt"`
}
