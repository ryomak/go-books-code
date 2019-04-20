package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", HandlerFunc)
	r.GET("/api", ApiHandler)
	r.GET("/", StaticHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func HandlerFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
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

func StaticHandler(c *gin.Context) {
	user := User{
		Name: "example",
		Age:  20,
	}
	c.HTML(http.StatusOK, "views/index.tmpl", gin.H{
		"Name": user.Name,
		"Age":  user.Age,
	})
}

//
// 応用 api
//

func ApiHandler(c *gin.Context) {
	user := User{
		Name: "example",
		Age:  20,
	}
	c.JSON(200, gin.H{
		"Name": user.Name,
		"Age":  user.Age,
	})
}
