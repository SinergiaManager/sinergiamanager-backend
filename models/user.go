package models

import "time"

type UserDb struct {
	ID       string    `bson:"_id"`
	Username string    `bson:"username"`
	Name     string    `bson:"name"`
	Surname  string    `bson:"surname"`
	Email    string    `bson:"email"`
	InsertAt time.Time `bson:"insert_at"`
	UpdateAt time.Time `bson:"update_at"`
}

type UserIns struct {
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	Surname  string    `json:"surname" bson:"surname"`
	Email    string    `json:"email" bson:"email"`
	InsertAt time.Time `json:"insertAt" bson:"insert_at"`
	UpdateAt time.Time `json:"updateAt" bson:"update_at"`
}
