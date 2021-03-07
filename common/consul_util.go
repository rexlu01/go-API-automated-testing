package common

import (
	"fmt"
	"go-api-automated-testing/config"

	consulapi "github.com/hashicorp/consul/api"
)

/*--------------------consul相关操作--------------------------*/
func GetConsulConn(config *config.Localserver) (*consulapi.Client, error) {
	// 连接consul
	conn := consulapi.DefaultConfig()
	conn.Address = config.ConsulAddress

	client, err := consulapi.NewClient(conn)
	if err != nil {
		fmt.Println("consul client error: ", err)
		return nil, err
	}
	return client, nil

}

func GetConsulKV(config *config.Localserver, key string) ([]byte, error) {
	client, err := GetConsulConn(config)
	if err != nil {
		panic(err)
	}

	data, _, err := client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}
	return data.Value, nil

}
