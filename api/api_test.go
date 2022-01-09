package api

import (
	"encoding/json"
	currency_rate "github.com/richguo0615/currency-rate/api/service/currency-rate"
	"net/http"
	"net/http/httptest"
	"testing"
)

// API - /currency/rate 測試
func TestCurrencyRates(t *testing.T) {
	srv := StartServer()

	type args struct {
		from   string //來源幣別
		to     string //目標幣別
		amount string //金額數字
	}
	type want struct {
		code int
		res  currency_rate.RateRes
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		// TODO: Add test cases.
		{
			name: "case1 - TWD -> USD 四捨五入",
			args: args{
				from:   "TWD",
				to:     "USD",
				amount: "1",
			},
			want: want{
				code: http.StatusOK,
				res: currency_rate.RateRes{
					Res: "0.03",
				},
			},
		},
		{
			name: "case2 - JPY -> USD 四捨五入",
			args: args{
				from:   "JPY",
				to:     "USD",
				amount: "1",
			},
			want: want{
				code: http.StatusOK,
				res: currency_rate.RateRes{
					Res: "0.01",
				},
			},
		},
		{
			name: "case3 - 未填from值",
			args: args{
				from:   "",
				to:     "JPY",
				amount: "1",
			},
			want: want{
				code: http.StatusBadRequest,
				res: currency_rate.RateRes{
					Error: "req.from can not be null",
				},
			},
		},
		{
			name: "case4 - twd -> JPY 千分位",
			args: args{
				from:   "twd",
				to:     "JPY",
				amount: "1000.12589",
			},
			want: want{
				code: http.StatusOK,
				res: currency_rate.RateRes{
					Res: "3,669.46",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest("GET", "/currency/rate", nil)
			q := req.URL.Query()
			q.Add("from", tt.args.from)
			q.Add("to", tt.args.to)
			q.Add("amount", tt.args.amount)
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)

			res := currency_rate.RateRes{}
			json.Unmarshal(w.Body.Bytes(), &res)

			if w.Code != tt.want.code {
				t.Errorf("want.code: %d, but got: %d", tt.want.code, w.Code)
			}
			switch w.Code {
			case http.StatusOK:
				if res.Res != tt.want.res.Res {
					t.Fatalf("want.Res: %s, but got: %s", tt.want.res.Res, res.Res)
				}
			case http.StatusBadRequest:
				if res.Error != tt.want.res.Error {
					t.Fatalf("want.Error: %s, but got: %s", tt.want.res.Error, res.Error)
				}
			}
		})
	}
}
