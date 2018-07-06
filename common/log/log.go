package log

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var logFile *os.File
var Logger *logging.Logger

func init() {
	var err error
	shiftLogs()
	Logger = logging.MustGetLogger("tempLog")
	var format = logging.MustStringFormatter(
		`[%{time:2006-01-02 15:04:05.000 CST}][%{level:.4s}](%{shortfile}) > %{message}`,
	)
	logFile, err = os.OpenFile("log/temp.log", os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("read logfile failed!")
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend2Formatter)
	//	Logger.Info("info")
	//	Logger.Notice("notice")
	//	Logger.Warning("warning")
	//	Logger.Error("xiaorui.cc")
	//	Logger.Critical("太严重了")
	fmt.Println("log init sucessed!")
}

func shiftLogs() {

	tempFile, err1 := os.OpenFile("log/temp.log", os.O_RDWR, os.ModePerm)
	if err1 != nil {
		fmt.Println("read temp.log failed")
	}
	primaryFile, err2 := os.OpenFile("log/primary.log", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err2 != nil {
		fmt.Println("read primary.log failed!")
	}
	defer tempFile.Close()
	defer primaryFile.Close()
	chunks := []byte{}
	buf := make([]byte, 1024)
	for {
		n, _ := tempFile.Read(buf)
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	primaryFile.WriteString("\n" + string(chunks))

	tempFile.Truncate(0)
}

func CloseLogFile() {
	logFile.Close()
}

//打印日志文档并输出控制台
func ErrorPrintToAll(args ...interface{}) {
	Logger.Error(args)
	fmt.Println(args)
}

//判断error,异常日志打印
func JudgeError(err error) bool {
	var orange bool
	if err != nil {
		Logger.Error("err:", err)
		orange = false
	} else {
		orange = true
	}
	return orange
}
