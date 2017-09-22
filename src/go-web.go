package main

import (
	"fmt"
	"os"
	//"flag"
	//"log"
	"webserver"
	"github.com/urfave/cli"
)

// command line flags
var (
	//port = flag.String("port", ":8080", "Listen port")
)

func main() {
	//flag.Parse()

	app := cli.NewApp()
	app.Name = "go-web"
	app.Usage = "Sample rest service app"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port, p",
			Value:  8081,
			EnvVar: "GOWEB_PORT",
			Usage:  "Listening port",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		err := validateContext(ctx)
		if err != nil {
			return err
		}

		port := fmt.Sprintf(":%s", ctx.String("port"))
		webserver := webserver.NewIrisWS()
		webserver.Listen(port)

		return nil
	}

	app.Run(os.Args)
}


func validateContext(ctx *cli.Context) error {
	if ctx.String("port") == "" {
		return fmt.Errorf("$GOWEB_PORT not set or not passed as a cli flag")
	}
	return nil
}