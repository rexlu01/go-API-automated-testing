package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api-automated-testing/common"
	"go-api-automated-testing/config"
	pb "go-api-automated-testing/data/server/proto"
)

type StoreData struct {
}

func (g *StoreData) Store(ctx context.Context, req *pb.StoreRequest, rsp *pb.StoreRespons) error {
	//这里写拿到处理好的数据往数据库里存
	testdata := []string{"1.1", "1.2"}

	req.Details = testdata
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
