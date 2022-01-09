package api

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	r := gin.Default()
	r.GET("/currency/rate", CurrencyRates)
	return r
}
