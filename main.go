package main

import (
	"github.com/richguo0615/currency-rate/api"
	_ "github.com/richguo0615/currency-rate/docs"
)

const (
	port = ":8880"
)

func main() {
	srv := api.StartServer()
	api.Swagger(srv, port)
	srv.Run(port)
}


