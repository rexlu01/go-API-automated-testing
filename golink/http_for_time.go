/*
按时间跑压测的实例,还有bug。
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var numCores = flag.Int("n", 2, "number of CPU cores to use")

type Promtest interface {
	SendRequest()
	RunTimeReq()
	ConcurrencyRunAndTotal() (float64, float64)
}

type RequestInfo struct {
	url           string
	method        string
	Concurrency   int64
	executionTime int64 `单位秒`
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
	fmt.Println("Finish")
}

func ConcurrencyRunAndTotal(reqinfo RequestInfo, finish chan bool, elapsedChin chan float64, reqnumChin chan bool, totalfinish chan bool, totalelapsedAVG chan float64, totalSuccessRate chan float64) {
	var totaleltime float64
	var successtreqotal int64
	totalNum := 0

	for i := 0; i < int(reqinfo.Concurrency); i++ {
		go RunTimeReq(reqinfo, finish, elapsedChin, reqnumChin)
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
	reqinfo.executionTime = 2
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
