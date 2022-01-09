package currency_rate

import (
	"github.com/richguo0615/currency-rate/api/model"
	"github.com/richguo0615/currency-rate/tools"
)

func CountRate(rateMap model.CurrencyRateMap, req RateReq) (float64, error) {
	rateList, _ := rateMap[req.From]
	rate, _ := rateList[req.To]

	res := tools.Round(req.Amount * rate, 2)

	return res, nil
}
