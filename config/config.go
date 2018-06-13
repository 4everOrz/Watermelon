package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	//运行配置
	DeBug      bool
	HandlePort string
	//redis设置
	RDhost     string
	RDport     string
	RDdb       int
	RDpassword string
	RDmaxidle  int
	RDmaxactiv int
	//mysql配置
	SqlType     string
	SqlHost     string
	SqlPort     string
	SqlDB       string
	SqlUser     string
	SqlPassword string
	//rabbitmq配置
	MQueueName    string
	MExchange     string
	MExchangeType string
	MMqurl        string
	MKey          string
}

func GetConf() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
}
