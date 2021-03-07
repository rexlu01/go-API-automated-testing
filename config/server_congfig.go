package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Localserver struct {
	ConsulAddress string
}

var ConsulAdd Localserver

func LoadJsonFile(filePath string) *Localserver {
	//打开json
	jsonFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &ConsulAdd)

	return &ConsulAdd

}
