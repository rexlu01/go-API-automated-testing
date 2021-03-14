package main

import (
	"context"
	pb "go-api-automated-testing/web/server/proto"
	"io"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type ShowData struct{}

var (
	cl  pb.ShowDataService
	Mut sync.Mutex
)

func WriteTo(w io.Writer) {
	//flag := c.Param("flag")

	response, err := cl.GetDataCrontab(context.TODO(), &pb.TryRequest{
		Flag: "1",
	})

	Mut.Lock()
	defer Mut.Unlock()

	cpuData := make(plotter.XYs, len(response.CPU))
	for i, p := range response.CPU {
		cpuData[i].X = float64(i + 1)
		cpuData[i].Y = p
	}

	memData := make(plotter.XYs, len(response.Mem))
	for i, p := range response.Mem {
		memData[i].X = float64(i + 1)
		memData[i].Y = p
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	cpuLine, err := plotter.NewLine(cpuData)
	if err != nil {
		log.Fatal(err)
	}
	cpuLine.Color = plotutil.Color(1)

	memLine, err := plotter.NewLine(memData)
	if err != nil {
		log.Fatal(err)
	}
	memLine.Color = plotutil.Color(2)

	p.Add(cpuLine, memLine)

	p.Legend.Add("cpu", cpuLine)
	p.Legend.Add("mem", memLine)

	p.X.Min = 0
	p.X.Max = float64(50)
	p.Y.Min = 0
	p.Y.Max = 100

	wc, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		log.Fatal(err)
	}
	wc.WriteTo(w)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func image(w http.ResponseWriter, r *http.Request) {
	WriteTo(w)
}

func (g *ShowData) GetDataCrontab(c *gin.Context) {
	log.Print("Received Say.Anything API request")

	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})

}

func main() {
	service := web.NewService(
		web.Name("go.micro.web.showdata"),
		web.RegisterTTL(time.Second*3),
		web.RegisterInterval(time.Second*3),
		web.Address(":8081"), //指定micro web端口号
	)

	service.Init()

	cl = pb.NewShowDataService("go.micro.srv.showdata", client.DefaultClient)

	router := gin.Default()
	router.GET("/", gin.WrapF(index))
	router.GET("/image", gin.WrapF(image))

	service.Handle("/router", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	//router.Run(":8082")
}
