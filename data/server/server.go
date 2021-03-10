package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api-automated-testing/common"
	"go-api-automated-testing/config"
	pb "go-api-automated-testing/data/server/proto"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/proto"
)

type StoreData struct {
}

func (g *StoreData) Store(ctx context.Context, req *pb.StoreRequest, rsp *pb.StoreRespons) error {
	//这里写拿到处理好的数据往数据库里存
	var d proto.Message
	details, err := ptypes.MarshalAny(d)
	if err != nil {
		panic(err)
	}

	req.Details = details
	return nil
}
func main() {

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
