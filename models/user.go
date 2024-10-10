package models

import "time"

type UserDb struct {
	ID       string    `bson:"_id"`
	Username string    `bson:"username"`
	Name     string    `bson:"name"`
	Surname  string    `bson:"surname"`
	Email    string    `bson:"email"`
	Password string    `bson:"password"`
	Role     string    `bson:"role"`
	InsertAt time.Time `bson:"insert_at"`
	UpdateAt time.Time `bson:"update_at"`
}

type UserIns struct {
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	Surname  string    `json:"surname" bson:"surname"`
	Email    string    `json:"email" bson:"email" `
	Password string    `json:"password" bson:"password"`
	Role     string    `json:"role" bson:"role" default:"user"`
	InsertAt time.Time `json:"insertAt" bson:"insert_at"`
	UpdateAt time.Time `json:"updateAt" bson:"update_at"`
}

type UserOut struct {
	ID       string    `json:"id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	Surname  string    `json:"surname" bson:"surname"`
	Email    string    `json:"email" bson:"email"`
	Role     string    `json:"role" bson:"role"`
	InsertAt time.Time `json:"insertAt" bson:"insert_at"`
	UpdateAt time.Time `json:"updateAt" bson:"update_at"`
}
