package warehouse

import "github.com/jackc/pgx/v5/pgxpool"

type Repository interface {
	GetAll()
	GetById()
	Create()
	Update()
	Delete()
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() {}

func (r *repository) GetById() {}

func (r *repository) Create() {}

func (r *repository) Update() {}

func (r *repository) Delete() {}
