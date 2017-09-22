package routes

import "gopkg.in/kataras/iris.v6"

//NotFound returns a custom error message when nothing was found
func NotFound(ctx *iris.Context) {
	ctx.JSON(iris.StatusNotFound, map[string]string{
		"message": "not found",
	})
}
