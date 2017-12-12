package log

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota + 1
	LevelInfo
	LevelWarn
	LevelError
)

var (
	logLevel = LevelDebug
)

func SetLevel(level LogLevel) {
	logLevel = level
}

func prefix() string {
	var ok bool
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	file = filepath.Base(file)
	return fmt.Sprintf("%s:%d", file, line)
}

func logEnable(level LogLevel) bool {
	return level >= logLevel
}

func Debug(args ...interface{}) {
	if logEnable(LevelDebug) {
		p := prefix()
		log.Println(p+" [D]", args)
	}
}

func Debugf(format string, args ...interface{}) {
	if logEnable(LevelDebug) {
		p := prefix()
		log.Println(p+" [D]", fmt.Sprintf(format, args))
	}
}

func Info(args ...interface{}) {
	if logEnable(LevelInfo) {
		p := prefix()
		log.Println(p+" [I]", args)
	}
}

func Infof(format string, args ...interface{}) {
	if logEnable(LevelInfo) {
		p := prefix()
		log.Println(p+" [I]", fmt.Sprintf(format, args))
	}
}

func Warn(args ...interface{}) {
	if logEnable(LevelWarn) {
		p := prefix()
		log.Println(p+" [W]:", args)
	}
}

func Warnf(format string, args ...interface{}) {
	if logEnable(LevelWarn) {
		p := prefix()
		log.Println(p+" [W]:", fmt.Sprintf(format, args))
	}
}

func Error(args ...interface{}) {
	if logEnable(LevelError) {
		p := prefix()
		log.Println(p+" [E]", args)
	}
}

func Errorf(format string, args ...interface{}) {
	if logEnable(LevelError) {
		p := prefix()
		log.Println(p+" [E]", fmt.Sprintf(format, args))
	}
}
