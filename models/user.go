package models

import (
	"regexp"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
)

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
	Username string    `json:"username" bson:"username" validate:"required"`
	Name     string    `json:"name" bson:"name" validate:"required"`
	Surname  string    `json:"surname" bson:"surname" validate:"required"`
	Email    string    `json:"email" bson:"email" validate:"required"`
	Password string    `json:"password" bson:"password" validate:"required"`
	Role     string    `json:"role" bson:"role"`
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

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(UserIns)

	/* input validation */
	if ok, _, _, _ := verifyPassword(user.Password); !ok {
		sl.ReportError(user.Password, "Password", "Password", "password", "")
	}

	match := `^[\w\-\.\d]+@[\w\-\.\d]+\.[\w\-\.\d]+$`
	if ok, _ := regexp.MatchString(match, user.Email); !ok {
		sl.ReportError(user.Email, "Email", "Email", "email", "")
	}
}

func verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false, false, false, false
		}
	}
	sevenOrMore = letters >= 8
	return
}
