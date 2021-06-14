package main

import (
	"context"
	"fmt"
	pbh "go-api-automated-testing/golink/server/proto"
	pb "go-api-automated-testing/web/server/proto"
	"html/template"
	"log"
	"net/http"

	"go-api-automated-testing/web/mon"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

type Fweb struct {
}

var (
	cl pb.FwebService
	//clh pbh.SendAPIService
)

var monit mon.Monitor

func (g *Fweb) SendMweb(c *gin.Context) {
	log.Print("Received Start.SendMessage API Request")
	name := c.Param("name")

	_, err := cl.MakeWeb(context.TODO(), &pb.FrontRequest{
		Name: name,
	})

	dataServer := pbh.NewSendAPIService("go.micro.srv.requestapi", client.DefaultClient)
	dataResp, err := dataServer.ProcessAPI(c, &pbh.SendRequest{RequestName: "ping/pong test", RequestURL: "http://47.115.20.3:81/ping", RequestMethod: "GET", IsPress: true, RunTime: 0, RunTimes: 2, Concurrency: 2})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataResp.ResponseTime)
	if err != nil {
		c.JSON(500, err)
	}

	c.HTML(200, "index.html", gin.H{"title": dataResp.ResponseTime})
}

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

	service := web.NewService(
		web.Name("go.micro.web.frontweb"),
		web.Address(":5977"),
	)

	service.Init()

	cl = pb.NewFwebService("go.micro.srv.frontweb", client.DefaultClient)

	mweb := new(Fweb)

	ginRouter := gin.Default()
	ginRouter.LoadHTMLGlob("html/*")
	////加载静态资源，例如网页的css、js
	ginRouter.Static("/static", "./static")

	ginRouter.GET("/sendmweb/:name", mweb.SendMweb)
	ginRouter.GET("/sendMstore")

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
