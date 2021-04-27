/*
按时间跑压测的实例,已修复。补充按次数跑案例
*/
package main

import (
	"context"
	"flag"
	"fmt"
	pb "go-api-automated-testing/golink/server/proto"
	"net/http"
	"time"
)

type SendAPI struct {
}

var numCores = flag.Int("n", 2, "number of CPU cores to use")

type Promtest interface {
	SendRequest()
	RunTimeReq()
	ConcurrencyRunAndTotal() (float64, float64)
}

type RequestInfo struct {
	url                string
	method             string
	Concurrency        int64
	executionType      int16 `1为执行时间，2为执行次数`
	executionfrequency int64 `执行次数, 如果执行type为1时，这个值为0`
	executionTime      int64 `单位秒，如果执行type为0时，这个值为0`
}

func SendRequest(reqinfo RequestInfo, elapsedChin chan float64, reqnumChin chan bool) {
	start := time.Now()
	resp, err := http.Get(reqinfo.url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	elapsed := time.Since(start).Seconds()
	elapsedChin <- elapsed

	if resp.StatusCode == 200 {
		reqnumChin <- true
	} else {
		reqnumChin <- false
	}
}

func RunNumsReq(reqinfo RequestInfo, finish chan bool, elapsedChin chan float64, reqnumChin chan bool) {
	loopchan := make(chan bool)
	var i int64
	go func() {
		for i = 0; i < reqinfo.executionfrequency; i++ {
			SendRequest(reqinfo, elapsedChin, reqnumChin)
		}
		loopchan <- true
	}()
	<-loopchan
	finish <- true
	fmt.Println("req Finish")
}

func RunTimeReq(reqinfo RequestInfo, finish chan bool, elapsedChin chan float64, reqnumChin chan bool) {
	timeout := time.After(time.Second * time.Duration(reqinfo.executionTime))
	loopchan := make(chan bool)
	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("timeout")
				loopchan <- true
				return
			default:
				SendRequest(reqinfo, elapsedChin, reqnumChin)
			}
		}
	}()
	<-loopchan
	finish <- true
	fmt.Println("req Finish")
}

func ConcurrencyRunAndTotal(reqinfo RequestInfo, finish chan bool, elapsedChin chan float64, reqnumChin chan bool, totalfinish chan bool, totalelapsedAVG chan float64, totalSuccessRate chan float64) {
	var totaleltime float64
	var successtreqotal int64
	totalNum := 0

	if reqinfo.executionType == 1 {
		for i := 0; i < int(reqinfo.Concurrency); i++ {
			go RunTimeReq(reqinfo, finish, elapsedChin, reqnumChin)
		}
	} else if reqinfo.executionType == 2 {
		for i := 0; i < int(reqinfo.Concurrency); i++ {
			go RunNumsReq(reqinfo, finish, elapsedChin, reqnumChin)
		}
	}

	go func() {
		for j := 0; j < int(reqinfo.Concurrency); j++ {
			<-finish
		}
		close(elapsedChin)
		close(reqnumChin)
	}()

	for eltime := range elapsedChin {
		totaleltime += eltime
	}
	totalNum = len(reqnumChin)
	for isscuess := range reqnumChin {
		if isscuess {
			successtreqotal++
		}
		//totalNum++
	}

	//计算平均响应时间
	elapsedAVG := totaleltime / float64(totalNum)

	//成功率
	SuccessRate := float64(successtreqotal) / float64(totalNum)

	totalelapsedAVG <- elapsedAVG
	totalSuccessRate <- SuccessRate

	totalfinish <- true
}

func main() {
	//初始化结构体
	var reqinfo RequestInfo
	//reqinfo.executionTime = 2
	//按次数试一把
	reqinfo.executionType = 2
	reqinfo.executionfrequency = 2000
	reqinfo.method = "GET"
	reqinfo.url = "http://47.115.20.3:81/ping"
	reqinfo.Concurrency = 10

	//定义协程
	finish := make(chan bool, reqinfo.Concurrency)
	elapsedChin := make(chan float64, 80000)
	reqnumChin := make(chan bool, 80000)
	totalfinish := make(chan bool)
	totalelapsedAVG := make(chan float64, 1)
	totalSuccessRate := make(chan float64, 1)
	go ConcurrencyRunAndTotal(reqinfo, finish, elapsedChin, reqnumChin, totalfinish, totalelapsedAVG, totalSuccessRate)

	go func() {
		for {
			<-totalfinish
		}
	}()

	elapsedAVG := <-totalelapsedAVG
	SuccessRate := <-totalSuccessRate

	fmt.Printf("%v\n", elapsedAVG) //单位s
	fmt.Printf("%v\n", SuccessRate)

}

func (g *SendAPI) ProcessAPI(ctx context.Context, req *pb.SendRequest, resp *pb.GetRespons) error {
	//初始化结构体
	var reqinfo RequestInfo
	reqinfo.url = req.RequestURL

	if req.IsPress {
		if req.RunTimes == "0" {
			reqinfo.executionType = 1
		} else if req.RunTime == "0" {
			reqinfo.executionType = 2
		}
	}

	req.RequestName = "ping/pong test"
	req.RequestURL = "http://47.115.20.3:81/ping"

	return nil

}
