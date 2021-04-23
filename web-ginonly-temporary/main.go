package main

import (
	"fmt"
	"go-api-automated-testing/web/mon"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var monit mon.Monitor

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/monitor.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func Image(w http.ResponseWriter, r *http.Request) {
	monit.WriteTo(w)
}

func main() {
	monit.MaxRecord = 50

	ginRouter := gin.Default()
	ginRouter.LoadHTMLGlob("html/*")
	////加载静态资源，例如网页的css、js
	ginRouter.Static("/static", "./static")
	ginRouter.GET("/sendmweb/", SendMweb)
	ginRouter.POST("/form", From)

	ginRouter.GET("/monit", gin.WrapF(Index))
	ginRouter.GET("/image", gin.WrapF(Image))

	go monit.Run()
	ginRouter.Run(":5656")
}

func SendMweb(c *gin.Context) {
	log.Print("Received Start.SendMessage API Request")

	c.HTML(200, "index.html", gin.H{})
}

func From(c *gin.Context) {

	userName := c.PostForm("requestname")

	c.String(http.StatusOK, fmt.Sprintf("username:%s", userName))

}
