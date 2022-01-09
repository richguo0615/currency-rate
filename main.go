package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richguo0615/currency-rate/api"
)

func main() {
	r := gin.Default()
	r.GET("/currency/rate", api.CurrencyRates)
	r.Run()
}


