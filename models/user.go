package models

import "time"

type User struct {
	username string    `json:"username"`
	name     string    `json:"name"`
	surname  string    `json:"surname"`
	email    string    `json:"email"`
	insertAt time.Time `json:"insert_at"`
	updateAt time.Time `json:"update_at"`
}
