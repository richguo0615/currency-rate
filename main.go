package main

import (
	"github.com/richguo0615/currency-rate/api"
)

func main() {
	srv := api.StartServer()
	srv.Run()
}


