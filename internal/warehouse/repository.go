package warehouse

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Warehouse, error)
	GetById(ctx context.Context)
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Warehouse, error) {
	sql := `SELECT id, name, code, created_at, updated_at 
			FROM warehouses ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var warehouses []Warehouse

	for rows.Next() {
		var w Warehouse
		err := rows.Scan(&w.Id, &w.Name, &w.Code, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}

func (r *repository) GetById(ctx context.Context) {}

func (r *repository) Create(ctx context.Context) {}

func (r *repository) Update(ctx context.Context) {}

func (r *repository) Delete(ctx context.Context) {}
