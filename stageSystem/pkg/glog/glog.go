package glog

import (
	"log"
	"os"
	"sync"
	"time"
)

const (
	LOG_FILE_NAME = "log.txt"
	YYYYMMDD      = "02.01.2006"
	HHMMSS24h     = "15:04:05"
)

type Logger struct {
	Filename   string
	FileHandle *os.File
	*log.Logger
}

var Glogger *Logger
var once sync.Once

func Init() *Logger {
	once.Do(func() {
		Glogger = createLogger(LOG_FILE_NAME)
	})
	return Glogger
}

func createLogger(fname string) *Logger {
	file, _ := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	ts := time.Now().Format(YYYYMMDD + " " + HHMMSS24h)

	return &Logger{
		Filename:   fname,
		FileHandle: file,
		Logger:     log.New(file, ts+" | ", log.Lshortfile),
	}
}
