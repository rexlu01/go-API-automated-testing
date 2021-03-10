package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api-automated-testing/common"
	"go-api-automated-testing/config"
	pb "go-api-automated-testing/data/server/proto"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/micro/go-micro/v2"
)

type StoreData struct {
}

func (g *StoreData) Store(ctx context.Context, req *pb.StoreRequest, rsp *pb.StoreRespons) error {
	//这里写拿到处理好的数据往数据库里存
	if req.DataBase == "mysql" {
		a := new(any.Any)
		a.Value = []byte("tps60")
		v := make([]*any.Any, 10)
		v = append(v, a)
		req.Details = v
	}
	rsp.DataId = 123101
	rsp.Message = "scess"
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("go_micro_srv_processdata"),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*3),
	)
	service.Init()

	// Register Handlers
	pb.RegisterStoreDataHandler(service.Server(), new(StoreData))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	localPath := common.GetLocalPath()
	pwd := localPath + "/config/server.json"
	fmt.Println(pwd)

	confserver := config.LoadJsonFile(pwd)

	value, err := common.GetConsulKV(confserver, "mysql")
	if err != nil {
		fmt.Println("loacl config fail")
	}
	var configMysql common.MysqlConfig
	if err := json.Unmarshal(value, &configMysql); err == nil {
		fmt.Println(configMysql.Charset)
	}

}
