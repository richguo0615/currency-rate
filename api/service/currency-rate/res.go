package currency_rate

type RateRes struct {
	Res     string `json: "res" example:"3.67"`
	Message string `json: "message" example:"invalid input"`
	Error   string `json: "error" example:"req.from can not be null"`
}
