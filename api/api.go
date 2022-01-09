package api

import (
	"github.com/gin-gonic/gin"
	rateSrv "github.com/richguo0615/currency-rate/api/service/currency-rate"
	"net/http"
)

var (
	rateJson = `{
	  "currencies": {
		"TWD": {
		  "TWD": 1,
		  "JPY": 3.669,
		  "USD": 0.03281
		},
		"JPY": {
		  "TWD": 0.26956,
		  "JPY": 1,
		  "USD": 0.00885
		},
		"USD": {
		  "TWD": 30.444,
		  "JPY": 111.801,
		  "USD": 1
		}
	  }
}`
)

func CurrencyRates(c *gin.Context) {
	data := rateSrv.ParseRateData(rateJson)
	rateMap, err := data.GetRateMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	req, err := rateSrv.GetReq(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error": err.Error(),
		})
		return
	}

	err = rateMap.Validate(req.From, req.To)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid input",
			"error": err.Error(),
		})
		return
	}

	res, err := rateSrv.CountRate(rateMap, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}
