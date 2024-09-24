package models

import "time"

type User struct {
	ID       string `bson:"_id"`
	Username string
	Name     string
	Surname  string
	Email    string
	InsertAt time.Time
	UpdateAt time.Time
}
