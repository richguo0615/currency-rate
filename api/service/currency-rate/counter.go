package currency_rate

import (
	"errors"
	"fmt"
	"github.com/richguo0615/currency-rate/api/model"
	"github.com/richguo0615/currency-rate/tools"
)

func GetRate(data model.CurrencyData, req RateReq) (float64, error) {
	from, to, amount := req.From, req.To, req.Amount

	rateMap, ok := data["currencies"]
	if !ok {
		return 0, errors.New("got wrong rateJson without .currencies")
	}

	rateList, ok := rateMap[from]
	if !ok {
		return 0, errors.New(fmt.Sprintf("rate data does not contains: %s", from))
	}

	rate, ok := rateList[to]
	if !ok {
		return 0, errors.New(fmt.Sprintf("rate data does not contains: %s.%s", from, to))
	}

	res := tools.Round(amount * rate, 2)

	return res, nil
}
