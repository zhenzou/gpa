package log

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

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

func Debug(args ...interface{}) {
	p := prefix()
	log.Println(p+" [D]", args)
}

func Info(args ...interface{}) {
	p := prefix()
	log.Println(p+" [I]", args)
}

func Infof(format string, args ...interface{}) {
	p := prefix()
	log.Println(p+" [I]", fmt.Sprintf(format, args))
}

func Debugf(format string, args ...interface{}) {
	p := prefix()
	log.Println(p+" [D]", fmt.Sprintf(format, args))
}

func Warn(args ...interface{}) {
	p := prefix()
	log.Println(p+" [W]:", args)
}

func Warnf(format string, args ...interface{}) {
	p := prefix()
	log.Println(p+" [W]:", fmt.Sprintf(format, args))
}

func Error(args ...interface{}) {
	p := prefix()
	log.Println(p+" [E]", args)
}

func Errorf(format string, args ...interface{}) {
	p := prefix()
	log.Println(p+" [E]", fmt.Sprintf(format, args))
}
