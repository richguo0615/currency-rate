### Start API server (option 1)
```
# mac
./app

# windows
./app.exe
```

### Start API server (option 2)
```
# install go 1.16
go mod tidy
go run main.go
```

### Swagger API Docs
> 開啟 [本地 Swagger API 文件](http://localhost:8880/swagger/index.html#/Currency)
> 說明：匯率換算API [GET] [{baseUrl}/currency/rate](http://localhost:8880/currency/rate)

### API 測試
```
go test -v -run=TestCurrencyRates ./api

# output
...
--- PASS: TestCurrencyRates (0.00s)
    --- PASS: TestCurrencyRates/case1_-_TWD_->_USD_四捨五入 (0.00s)
    --- PASS: TestCurrencyRates/case2_-_JPY_->_USD_四捨五入 (0.00s)
    --- PASS: TestCurrencyRates/case3_-_未填from值 (0.00s)
    --- PASS: TestCurrencyRates/case4_-_twd_->_JPY_千分位 (0.00s)
PASS
```


