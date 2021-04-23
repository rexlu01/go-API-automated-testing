package main

import (
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
	go monit.Run()

	ginRouter := gin.Default()
	ginRouter.LoadHTMLGlob("html/*")
	////加载静态资源，例如网页的css、js
	ginRouter.Static("/static", "./static")
	ginRouter.GET("/sendmweb/", SendMweb)

	ginRouter.POST("/form", func(c *gin.Context) {

		_ = c.PostForm("requestname")

		c.Request.URL.Path = "/monit"
		c.Request.Method = "GET"
		ginRouter.HandleContext(c)

	})

	ginRouter.GET("/monit", gin.WrapF(Index))
	ginRouter.GET("/image", gin.WrapF(Image))

	ginRouter.Run(":5656")
}

func SendMweb(c *gin.Context) {
	log.Print("Received Start.SendMessage API Request")

	c.HTML(200, "index.html", gin.H{})
}
