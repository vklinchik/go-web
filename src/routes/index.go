package routes

import "gopkg.in/kataras/iris.v6"

func Index(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "<h1> Welcome to go-web!</h1>")
}
