package currency_rate

type RateRes struct {
	Result  string `json: "result"`
	Message string `json: "message"`
	Error   string `json: "error"`
}
