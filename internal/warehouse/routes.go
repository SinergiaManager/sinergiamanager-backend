package warehouse

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kataras/iris/v12"
)

type Module struct {
	handler *Handler
}

func NewModule(db *pgxpool.Pool) *Module {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{handler: handler}
}

func (m *Module) RegisterRoutes(party iris.Party) {
	warehouses := party.Party("/warehouses")
	{
		warehouses.Get("/", m.handler.GetAll)
		warehouses.Get("/{id: int}", m.handler.GetById)
		warehouses.Post("/", m.handler.Create)
		warehouses.Put("/{id:int}", m.handler.Update)
		warehouses.Delete("/{id:int}", m.handler.Delete)
	}

}
