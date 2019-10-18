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

func (logger *Logger) Debug(msg interface{}) (n int, err error) {
	logger.err_level = DEBUG
	return logger.Write([]byte(logger.convert(msg)))
}

func (logger *Logger) Notice(msg interface{}) (n int, err error) {
	logger.err_level = NOTICE
	n, err = logger.Write([]byte(logger.convert(msg)))
	return n, err
}

func (logger *Logger) Warning(msg interface{}) (n int, err error) {
	logger.err_level = WARNING
	n, err = logger.Write([]byte(logger.convert(msg)))
	return n, err
}

func (logger *Logger) Error(msg interface{}) (n int, err error) {
	logger.err_level = ERROR
	n, err = logger.Write([]byte(logger.convert(msg)))
	return n, err
}

func (logger *Logger) Write(msg []byte) (n int, err error) {
	if logger.Device == nil {
		panic("Write log msg to device error. check device setup.")
	}
	if logger.Level == 0 {
		logger.Level = ERROR | NOTICE | DEBUG | WARNING
	}
	if logger.Level&logger.err_level <= 0 {
		return
	}
	var levelStr string = logger.levelToString(logger.err_level)
	msg = []byte(fmt.Sprintf("[%s] %s %s\n", levelStr, time.Now().Format("2006-01-01 15:04:05"), msg))
	n, err = logger.Device.Write(msg)
	return n, err
}

func (logger *Logger) convert(msg interface{}) string {
	return fmt.Sprintf("%s", msg)
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
