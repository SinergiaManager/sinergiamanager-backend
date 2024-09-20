package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

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
