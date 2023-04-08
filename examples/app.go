package main

import (
	"github.com/go-labx/lightning"
	"github.com/lightning-contrib/static"
)

func main() {
	app := lightning.DefaultApp()

	app.Use(static.New(
		static.WithRoot("./public"),
		static.WithPrefix("/static"),
	))

	app.Get("/ping", func(ctx *lightning.Context) {
		ctx.Text(200, "hello world")
	})

	app.Run()
}
