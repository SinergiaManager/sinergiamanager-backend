package models

type Login struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type Register struct {
	Email           string `json:"email" bson:"email" validate:"required,email"`
	Password        string `json:"password" bson:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password" validate:"required,eqfield=Password"`
	Username        string `json:"username" bson:"username" validate:"required"`
	Name            string `json:"name" bson:"name" validate:"required"`
	Surname         string `json:"surname" bson:"surname" validate:"required"`
}
