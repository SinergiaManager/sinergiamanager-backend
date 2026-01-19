package warehouse

type CreateWarehouse struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code,omitempty"`
}

type UpdateWarehouse struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}
