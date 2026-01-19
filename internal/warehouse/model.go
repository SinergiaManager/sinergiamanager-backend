package warehouse

import "time"

type Warehouse struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Code      *string   `json:"code" db:"code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
