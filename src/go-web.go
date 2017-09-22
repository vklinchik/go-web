package main

import (
	"fmt"
	//"flag"
	//"log"
	"webserver"
)

// command line flags
var (
	//port = flag.String("port", ":8080", "Listen port")
)

func main() {
	//flag.Parse()

	webserver := webserver.NewIrisWS()
	webserver.Listen(":8080")

	fmt.Scanln() // keep it running TODO: revisit this
}