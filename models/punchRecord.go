package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type PunchRecordDb struct {
	ID         string    `bson:"_id"`
	EmployeeID string    `bson:"employee_id"`
	PunchType  string    `bson:"punch_type"`
	PunchTime  time.Time `bson:"punch_time"`
	Location   string    `bson:"location"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

type PunchRecordIns struct {
	EmployeeID string    `json:"employee_id" bson:"employee_id" validate:"required"`
	PunchType  string    `json:"punch_type" bson:"punch_type" validate:"required"`
	PunchTime  time.Time `json:"punch_time" bson:"punch_time"`
	Location   string    `json:"location" bson:"location" validate:"required"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

type PunchRecordOut struct {
	ID         string    `json:"id" bson:"_id"`
	EmployeeID string    `json:"employee_id" bson:"employee_id"`
	PunchType  string    `json:"punch_type" bson:"punch_type"`
	PunchTime  time.Time `json:"punch_time" bson:"punch_time"`
	Location   string    `json:"location" bson:"location"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

func PunchRecordStructLevelValidation(sl validator.StructLevel) {
	punchRecord := sl.Current().Interface().(PunchRecordIns)

	if punchRecord.PunchType != "IN" && punchRecord.PunchType != "OUT" {
		sl.ReportError(punchRecord.PunchType, "PunchType", "PunchType", "punchType", "")
	}

	if punchRecord.PunchTime.IsZero() {
		punchRecord.PunchTime = time.Now()
	}
}
