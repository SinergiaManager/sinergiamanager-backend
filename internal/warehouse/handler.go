package warehouse

import "github.com/kataras/iris/v12"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(ctx iris.Context) {}

func (h *Handler) GetById(ctx iris.Context) {}

func (h *Handler) Create(ctx iris.Context) {}

func (h *Handler) Update(ctx iris.Context) {}

func (h *Handler) Delete(ctx iris.Context) {}
