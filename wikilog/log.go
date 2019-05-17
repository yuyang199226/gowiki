package wikilog

import (
	"log"
	"os"
	"fmt"
	"io"
)

var (
	outLogger *log.Logger
	errLogger *log.Logger
)
func init() {
	errFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("打开日志文件失败", err)
	}
	outLogger = log.New(os.Stdout, "[WIKI]", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger = log.New(io.MultiWriter(os.Stderr, errFile), "[WIKI]", log.Ldate|log.Ltime|log.Lshortfile)

}


func Info(v ...interface{}) {
	outLogger.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	outLogger.Output(2, fmt.Sprintln(v...))
}