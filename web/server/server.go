package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/micro/go-micro"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	pb "go-api-automated-testing/web/server/proto"
)

type ShowData struct{}

// var (
// 	Mem []float64
// 	CPU []float64
// )

// func Run() {
// 	for {
// 		Collect()
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func Collect() ([]float64, []float64) {
	var Mem []float64
	var CPU []float64

	mem, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	cpu, err := cpu.Percent(500*time.Millisecond, false)
	if err != nil {
		log.Fatal(err)
	}

	//这里是个互斥锁
	(*sync.Mutex).Lock.Lock()
	defer (*sync.Mutex).Lock.Unlock()

	Mem = append(Mem, mem.UsedPercent)
	CPU = append(CPU, cpu[0])
	return Mem, CPU
}

func (g *ShowData) GetDataCrontab(ctx context.Context, req *pb.TryRequest, rsp *pb.TryRespons) error {
	if req.Flag == 1 {
		rsp.Mem, rsp.CPU = Collect()
	} else {
		rsp.Mem, rsp.CPU = nil, nil
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.showdata"),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*3),
	)
	service.Init()
	pb.RegisterShowDataHandler(service.Server(), new(ShowData))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
