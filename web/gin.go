package main

import (
	pb "go-api-automated-testing/web/server/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

type Mweb struct{}

var (
	cl pb.MwebService
)

func (g *Mweb) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func main() {
	service := web.NewService(
		web.Name("go.micro.web.sendmweb"),
	)

	service.Init()

	cl = pb.NewMwebService("go.micro.srv.sendmweb", client.DefaultClient)

	mweb := new(Mweb)

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", mweb.Anything)
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	//注册handler
	service.Handle("/", router)
	//run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
