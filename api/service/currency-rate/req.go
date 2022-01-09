package currency_rate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RateReq struct {
	From   string
	To     string
	Amount float64
}

func GetReq(c *gin.Context) (req RateReq, err error) {
	from, to := c.Query("from"), c.Query("to")
	amountStr := c.Query("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return RateReq{}, errors.New("amount is required float type")
	}

	return RateReq{
		From:   from,
		To:     to,
		Amount: amount,
	}, nil
}
