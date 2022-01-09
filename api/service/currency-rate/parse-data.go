package currency_rate

import (
	"encoding/json"
	"github.com/richguo0615/currency-rate/api/model"
)

func ParseRateData(rateJson string) (d model.CurrencyData) {
	_ = json.Unmarshal([]byte(rateJson), &d)
	return d
}
