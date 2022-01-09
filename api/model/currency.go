package model

type CurrencyData map[string]CurrencyRateMap

type CurrencyRateMap map[string]Rate

type Rate map[string]float64
