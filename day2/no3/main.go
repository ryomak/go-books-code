package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://api.bitflyer.jp/v1/getticker?product_code="

type Coin struct {
	Name  string
	Query string
}

type Response struct {
	ProductCode string  `json:"product_code"`
	LTP         float64 `json:"ltp"`
}

func main() {
	coin := Coin{Name: "Bitcoin", Query: "btc_jpy"}
	resp, err := http.Get(baseUrl + coin.Query)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	mapRes := new(Response)
	if err := json.Unmarshal(byteArray, &mapRes); err != nil {
		panic(err)
	}
	fmt.Printf("プロダクト名：%+v\n", mapRes.ProductCode)
	fmt.Printf("価格：%+v\n", mapRes.LTP)
}
