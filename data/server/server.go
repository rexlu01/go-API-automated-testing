package main

import (
	"encoding/json"
	"fmt"
	"go-api-automated-testing/common"
	"go-api-automated-testing/config"
)

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
