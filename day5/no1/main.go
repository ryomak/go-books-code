package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", HandlerFunc)
	http.HandleFunc("/api", ApiHandler)
	http.HandleFunc("/", StaticHandler)

	fmt.Println("server start port:8080")
	http.ListenAndServe(":8080", nil)
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

//
//応用→静的ページ(template)をレンダリング　APIサーバ作成
//

type User struct {
	Name string
	Age  int
}

//
// 応用 template
//

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "example",
		Age:  20,
	}
	tmpl := template.Must(template.ParseFiles("./views/index.tmpl"))
	tmpl.Execute(w, user)
}

//
// 応用 api
//

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "example",
		Age:  20,
	}
	// simple health check
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
