package model

import (
	"errors"
	"fmt"
)

type CurrencyData map[string]CurrencyRateMap

type CurrencyRateMap map[string]Rate

type Rate map[string]float64

func (d CurrencyData) GetRateMap() (CurrencyRateMap, error) {
	rateMap, ok := d["currencies"]
	if !ok {
		return nil, errors.New("got wrong rateJson without .currencies")
	}
	return rateMap, nil
}

func (m CurrencyRateMap) Validate(from, to, amount string) error {
	if from == "" {
		return errors.New(fmt.Sprintf("req.from can not be null"))
	}
	if to == "" {
		return errors.New(fmt.Sprintf("req.to can not be null"))
	}
	if amount == "" {
		return errors.New(fmt.Sprintf("req.amount can not be null"))
	}

	rateList, ok := m[from]
	if !ok {
		return errors.New(fmt.Sprintf("req.from can not be: %s", from))
	}

	_, ok = rateList[to]
	if !ok {
		return errors.New(fmt.Sprintf("req.to can not be: %s", to))
	}

	return nil
}