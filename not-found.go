package main

import (
	"github.com/kataras/iris/v12"
)

type notFoundResponse struct {
	Message string   `json:"message"`
	Routes  []string `json:"routes,omitempty"`
}

func notFound(ctx iris.Context) {
	suggestPaths := ctx.FindClosest(3)

	if len(suggestPaths) == 0 {
		ctx.JSON(notFoundResponse{
			Message: "Route not found",
		})
		return
	}

	ctx.JSON(notFoundResponse{
		Message: "Route not found. Check these similar routes and try again",
		Routes:  suggestPaths,
	})
}
