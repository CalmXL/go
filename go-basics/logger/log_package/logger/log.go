package logger

import (
	"io"
	"log"
	"os"
)

const flag = log.Ldate | log.Ltime | log.Lshortfile

var (
	logFile     io.Writer
	logFilePath = ". /logs/global.log"
)

const (
	infoPrefix  = "[INFO] "
	warnPrefix  = "[WARN] "
	errorPrefix = "[ERROR] "
	debugPrefix = "[DEBUG] "
)

var (
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
	debugLog *log.Logger
)

func New() error {
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	infoLog = log.New(logFile, infoPrefix, flag)
	warnLog = log.New(logFile, warnPrefix, flag)
	errorLog = log.New(logFile, errorPrefix, flag)
	debugLog = log.New(logFile, debugPrefix, flag)

	return nil
}

func Info(format string, vrs ...interface{}) {
	infoLog.Printf(format, vrs...)
}

func Warn(format string, vrs ...interface{}) {
	warnLog.Printf(format, vrs...)
}

func Error(format string, vrs ...interface{}) {
	errorLog.Printf(format, vrs...)
}

func Debug(format string, vrs ...interface{}) {
	debugLog.Printf(format, vrs...)
}
