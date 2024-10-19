package models

import (
	"time"
)

type ConfigDb struct {
	ID           string    `bson:"_id"`
	SupportEmail string    `bson:"support_email"`
	SmtpHost     string    `bson:"smtp_host"`
	SmtpPort     int       `bson:"smtp_port"`
	SmtpEmail    string    `bson:"smtp_user"`
	SmtpPassword string    `bson:"smtp_pass"`
	InsertAt     time.Time `bson:"insert_at"`
	UpdateAt     time.Time `bson:"update_at"`
}

type ConfigIns struct {
	SupportEmail string    `json:"support_email" bson:"support_email" validate:"required,email"`
	SmtpHost     string    `json:"smtp_host" bson:"smtp_host" validate:"required,hostname"`
	SmtpPort     int       `json:"smtp_port" bson:"smtp_port" validate:"required,hostname_port"`
	SmtpEmail    string    `json:"smtp_user" bson:"smtp_user" validate:"required,email"`
	SmtpPassword string    `json:"smtp_pass" bson:"smtp_pass" validate:"required"`
	InsertAt     time.Time `json:"insert_at" bson:"insert_at"`
	UpdateAt     time.Time `json:"update_at" bson:"update_at"`
}
