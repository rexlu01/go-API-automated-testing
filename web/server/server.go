package main

import (
	"context"
	pb "go-api-automated-testing/web/server/proto"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
)

type Fweb struct {
}

func (g *Fweb) MakeWeb(ctx context.Context, req *pb.FrontRequest, rsp *pb.FrontRespons) error {
	rsp.Msg = "this is response " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.frontweb"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	pb.RegisterFwebHandler(service.Server(), new(Fweb))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
