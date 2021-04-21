package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	ginRouter := gin.Default()
	ginRouter.LoadHTMLGlob("html/*")
	////加载静态资源，例如网页的css、js
	ginRouter.Static("/static", "./static")
	ginRouter.GET("/sendmweb/", SendMweb)
	ginRouter.POST("/form", From)
	ginRouter.Run(":5656")
}

func SendMweb(c *gin.Context) {
	log.Print("Received Start.SendMessage API Request")
	//name := c.Param("name")

	// response, err := cl.MakeWeb(context.TODO(), &pb.FrontRequest{
	// 	Name: name,
	// })

	// if err != nil {
	// 	c.JSON(500, err)
	// }

	//c.HTML(200, "index.html", gin.H{"title": response})

	c.HTML(200, "index.html", gin.H{})
}

func From(c *gin.Context) {

	userName := c.PostForm("requestname")

	c.String(http.StatusOK, fmt.Sprintf("username:%s", userName))

}
