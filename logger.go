package logger

import (
	"fmt"
	"io"
	"time"
)

const (
	ERROR int = 1 << iota
	WARNING
	NOTICE
	DEBUG
)

type Logger struct {
	Device    io.Writer
	Level     int
	err_level int
}

func (logger *Logger) Debug(msg string) (n int, err error) {
	log.err_level = DEBUG
	return log.Write([]byte(msg))
}

func (logger *Logger) Notice(msg string) (n int, err error) {
	log.err_level = NOTICE
	n, err = log.Write([]byte(msg))
	return n, err
}

func (logger *Logger) Warning(msg string) (n int, err error) {
	log.err_level = WARNING
	n, err = log.Write([]byte(msg))
	return n, err
}

func (logger *Logger) Error(msg string) (n int, err error) {
	log.err_level = ERROR
	n, err = log.Write([]byte(msg))
	return n, err
}

func (logger *Logger) Write(msg []byte) (n int, err error) {
	if log.Device == nil {
		panic("Write log msg to device error. check device setup.")
	}
	if log.Level == 0 {
		log.Level = ERROR | NOTICE | DEBUG | WARNING
	}
	//fmt.Println(logger.err_level, log.Level&log.err_level)
	if log.Level&log.err_level <= 0 {
		return
	}
	var levelStr string = log.levelToString(logger.err_level)
	msg = []byte(fmt.Sprintf("[%s] %s %s\n", levelStr, time.Now().Format("2006-01-01 15:04:05"), msg))
	n, err = log.Device.Write(msg)
	return n, err
}

func (logger *Logger) levelToString(err_level int) string {
	levelMap := map[int]string{
		ERROR:   "Error",
		WARNING: "Warning",
		NOTICE:  "Notice",
		DEBUG:   "Debug",
	}
	return levelMap[err_level]
}
