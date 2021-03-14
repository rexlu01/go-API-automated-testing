package main

import (
	"context"
	pb "go-api-automated-testing/web/server/proto"
	"html/template"
	"log"
	"net/http"

	"go-api-automated-testing/web/mon"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
)

type Fweb struct {
}

var (
	cl pb.FwebService
)

var monit mon.Monitor

func (g *Fweb) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (g *Fweb) SendMweb(c *gin.Context) {
	log.Print("Received Start.SendMessage API Request")
	name := c.Param("name")

	response, err := cl.MakeWeb(context.TODO(), &pb.FrontRequest{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.HTML(200, "login.html", gin.H{"title": response})

}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/login.html")
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

	service := web.NewService(
		web.Name("go.micro.web.frontweb"),
		web.Address(":5977"),
	)

	service.Init()

	cl = pb.NewFwebService("go.micro.srv.frontweb", client.DefaultClient)

	mweb := new(Fweb)

	ginRouter := gin.Default()
	ginRouter.LoadHTMLGlob("html/*")

	ginRouter.GET("/sendmweb", mweb.Anything)
	ginRouter.GET("/sendmweb/:name", mweb.SendMweb)

	ginRouter.GET("/monit", gin.WrapF(Index))
	ginRouter.GET("/image", gin.WrapF(Image))

	go monit.Run()

	//注册handler
	service.Handle("/", ginRouter)

	//run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

// func main() {

// 	router := gin.Default()
// 	router.GET("/", gin.WrapF(index))
// 	router.GET("/image", gin.WrapF(image))

// 	go monitor.Run()

// 	router.Run(":8083")
// }
