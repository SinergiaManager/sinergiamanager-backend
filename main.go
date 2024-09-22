package main

import (
	Config "github.com/SinergiaManager/sinergiamanager-backend/config"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	Config.ConnectDb()

	type PingResponse struct {
		Message string `json:"message"`
	}

	app.Get("/ping", func(ctx iris.Context) {
		res := PingResponse{
			Message: "pong",
		}

		ctx.JSON(res)
	})

	app.Listen(":8080")
}
