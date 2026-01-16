package warehouse

import "time"

type Warehouse struct {
	Id       int
	Name     string
	Code     string
	UpdateAt time.Time
	InsertAt time.Time
}
