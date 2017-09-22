package webserver


import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"routes"
)

func NewIrisWS() *iris.Framework {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	app.OnError(404, routes.NotFound)

	app.Get("/", routes.Index) // root

	api := app.Party("/go-web/v1.0/demo-web")
	{
		api.Get("/", routes.Index)
		api.Get("/user/:id", routes.GetUser())
		api.Post("/user", routes.AddUser())
		/*
		api.Post("/snapshot", routes.CreateSnapshot(producer, cliCtx, srCtx))
		api.Get("/snapshot", routes.GetSnapshotInfo(producer, cliCtx))
		api.Post("/bootstrap", routes.Bootstrap(producer, cliCtx, srCtx))
		api.Get("/latest_offset", routes.LatestOffset(producer))
		*/
	}

	return app
}
