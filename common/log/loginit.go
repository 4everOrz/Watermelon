package log

import (
	"fmt"

	log4 "github.com/alecthomas/log4go"
)

func init() {
	log4.LoadConfiguration("log/log_conf.xml")
	fmt.Println("log init sucessed!")
}
