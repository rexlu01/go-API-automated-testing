package main

import (
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
	fmt.Printf("kv :%v", value)

}
