package config

import (
	"fmt"
	"strconv"

	"github.com/kylelemons/go-gypsy/yaml"
)

var ConfigFile *yaml.File

func init() {
	var err error
	ConfigFile, err = yaml.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("read config file failed!")
	}
	fmt.Println("config init successed!")
}
func GetString(key string) string {
	str, err := ConfigFile.Get(key)
	if err != nil {
		fmt.Println("Get  value of " + key + " failed")
	}
	return str
}
func GetInt(key string) int {
	str, err := ConfigFile.Get(key)
	if err != nil {
		fmt.Println("Get  value of " + key + " failed")
	}
	IntS, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("convert the key " + key + " to int failed ")
	}
	return IntS
}
func GetInt64(key string) int64 {
	Int64S, err := ConfigFile.GetInt(key)
	if err != nil {
		fmt.Println("Get  value of " + key + " failed")
	}
	return Int64S
}
func GetBool(key string) bool {
	bools, err := ConfigFile.GetBool(key)
	if err != nil {
		fmt.Println("Get  value of " + key + " failed")
	}
	return bools
}
