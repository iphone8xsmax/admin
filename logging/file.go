package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var(
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}


func openLogfile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):  //文件或目录是否存在
		mkDir(getLogFilePath())
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)  //权限是否满足
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil{
		log.Fatalf("Failed to openfile: %v", err)
	}
	return handle
}

//创建日志文件
func mkDir(string) {
	dir, _ := os.Getwd() //返回当前目录的路径
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil{
		panic(err)
	}
}
