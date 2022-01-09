package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rateSrv "github.com/richguo0615/currency-rate/api/service/currency-rate"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"

	_ "github.com/richguo0615/currency-rate/docs"
)

var rateJson = `{
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

// @BasePath /

// RateExample godoc
// @Summary 計算匯率
// @Schemes
// @Param   from    query    string     true    "來源幣別" Enums(TWD, JPY, USD)
// @Param   to     	query    string     true    "目標幣別" Enums(TWD, JPY, USD)
// @Param   amount  query    number		true    "金額數字" default(0)
// @Description 填入 "來源幣別"、"目標幣別"、"金額數字"，計算出當前匯率結果。
// @Tags Currency
// @Accept json
// @Produce json
// @Success 200 {object} currency_rate.RateRes "{"res": "3.67"}"
// @Success 400 {object} currency_rate.RateRes "{"message": "invalid input", "error": "req.from can not be null"}"
// @Router /currency/rate [get]
func CurrencyRates(c *gin.Context) {
	data := rateSrv.ParseRateData(rateJson)
	rateMap, err := data.GetRateMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, rateSrv.RateRes{
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	from, to, amount := getRateQuery(c)
	err = rateMap.Validate(from, to, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, rateSrv.RateRes{
			Message: "invalid input",
			Error:   err.Error(),
		})
		return
	}

	req, err := rateSrv.GetReq(from, to, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, rateSrv.RateRes{
			Message: "invalid input",
			Error:   err.Error(),
		})
		return
	}

	res, err := rateSrv.CountRate(rateMap, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, rateSrv.RateRes{
			Message: "internal server error",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rateSrv.RateRes{
		Res:     res,
		Message: "success",
		Error:   "",
	})
}

func getRateQuery(c *gin.Context) (from, to, amount string) {
	return strings.ToUpper(c.Query("from")), strings.ToUpper(c.Query("to")), strings.ToUpper(c.Query("amount"))
}

func Swagger(r *gin.Engine, port string) {
	url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", port))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}