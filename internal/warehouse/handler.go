package warehouse

import "github.com/kataras/iris/v12"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(ctx iris.Context) {
	warehouses, err := h.service.GetAll(ctx.Request().Context())
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"message": "Failed to fetch warehouses",
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": warehouses})
}

func (h *Handler) GetById(ctx iris.Context) {}

func (h *Handler) Create(ctx iris.Context) {}

func (h *Handler) Update(ctx iris.Context) {}

func (h *Handler) Delete(ctx iris.Context) {}
