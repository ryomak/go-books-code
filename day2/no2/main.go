package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://api.bitflyer.jp/v1/getticker?product_code="

type Coin struct {
	Name  string
	Query string
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
	fmt.Println(string(byteArray))
}
