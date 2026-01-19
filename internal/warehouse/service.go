package warehouse

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]Warehouse, error)
	GetById(ctx context.Context)
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(ctx context.Context) ([]Warehouse, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) GetById(ctx context.Context) {}

func (s *service) Create(ctx context.Context) {}

func (s *service) Update(ctx context.Context) {}

func (s *service) Delete(ctx context.Context) {}
