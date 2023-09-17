package main

import (
	"github.com/kataras/iris/v12"
)

func notFound(ctx iris.Context) {
	suggestPaths := ctx.FindClosest(3)

	if len(suggestPaths) == 0 {
		ctx.JSON(struct {
			Message string `json:"message"`
		}{Message: "Route not found"})
		return
	}

	ctx.JSON(struct {
		Message string   `json:"message"`
		Routes  []string `json:"routes"`
	}{
		Message: "Route not found. Check these similar routes and try again",
		Routes:  suggestPaths,
	})
}
