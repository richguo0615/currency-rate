package currency_rate

import (
	"errors"
	"strconv"
	"strings"
)

type RateReq struct {
	From   string
	To     string
	Amount float64
}

func GetReq(from, to, amountStr string) (req RateReq, err error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return RateReq{}, errors.New("req.amount is required float type")
	}

	return RateReq{
		From:   strings.ToUpper(from),
		To:     strings.ToUpper(to),
		Amount: amount,
	}, nil
}
