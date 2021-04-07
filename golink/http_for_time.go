/*
按时间跑压测的实例,有bug.
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

	if resp.StatusCode == 200 {
		elapsedChin <- elapsed
		reqnumChin <- true
		return
	}
	elapsedChin <- elapsed
	reqnumChin <- false
	return
}

func RunTimeReq(reqinfo RequestInfo, finish chan bool, elapsedChin chan float64, reqnumChin chan bool) {
	timeout := time.After(time.Second * time.Duration(reqinfo.executionTime))
	//finish = make(chan bool, reqinfo.Concurrency)

	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("timeout")
				finish <- true
				return
			default:
				SendRequest(reqinfo, elapsedChin, reqnumChin)
			}
		}
	}()

	finish <- true

}

func ConcurrencyRunAndTotal(reqinfo RequestInfo, finish chan bool, elapsedChin chan float64, reqnumChin chan bool, totalfinish chan bool) (float64, float64) {
	var totaleltime float64
	var successtreqotal int64

	for i := 0; i < int(reqinfo.Concurrency); i++ {
		go RunTimeReq(reqinfo, finish, elapsedChin, reqnumChin)
	}

	go func() {
		for {
			<-finish
		}
	}()
	close(elapsedChin)
	close(reqnumChin)
	totalNum := len(reqnumChin)

	for eltime := range elapsedChin {
		totaleltime += eltime
	}

	for isscuess := range reqnumChin {
		if isscuess == true {
			successtreqotal++
		}
	}

	//计算平均响应时间
	elapsedAVG := totaleltime / float64(totalNum)

	//成功率
	SuccessRate := float64(successtreqotal) / float64(totalNum)

	totalfinish <- true

	return elapsedAVG, SuccessRate

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
	elapsedAVG, SuccessRate := ConcurrencyRunAndTotal(reqinfo, finish, elapsedChin, reqnumChin, totalfinish)

	go func() {
		for {
			<-totalfinish
		}
		//close(totalfinish)
	}()

	fmt.Println(elapsedAVG)
	fmt.Println()
	fmt.Println(SuccessRate)

}
