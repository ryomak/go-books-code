package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://www.google.com/"

func main() {
	resp, err := http.Get(baseUrl)
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
