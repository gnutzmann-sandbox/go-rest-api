package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"go-rest-api/books"
	"go-rest-api/config"
)

func main() {
	config.LoadConfig()

	app := iris.New()
	app.Use(iris.Compression)
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.UseRouter(recover.New())

	books.ServeRoutes(app)

	err := app.Listen(":" + config.Config.Port)
	if err != nil {
		panic("Error to start server")
	}
}
