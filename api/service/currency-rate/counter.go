package currency_rate

import (
	"github.com/richguo0615/currency-rate/api/model"
	"github.com/richguo0615/currency-rate/tools"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CountRate(rateMap model.CurrencyRateMap, req RateReq) (string, error) {
	rateList, _ := rateMap[req.From]
	rate, _ := rateList[req.To]

	// 四捨五入
	res := tools.Round(req.Amount * rate, 2)

	// 千分位
	p := message.NewPrinter(language.English)

	return p.Sprintf("%.2f", res), nil
}
