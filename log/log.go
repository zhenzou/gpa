// +build !prod

package log

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var (
	logLevel = LevelDebug
)

func SetLevel(level level) {
	logLevel = level
}

func logEnable(level level) bool {
	return level >= logLevel
}

func Debug(args ...interface{}) {
	if logEnable(LevelDebug) {
		log.SetPrefix(LevelDebug.String())
		log.Output(3, fmt.Sprintln(args...))
	}
}

func Debugf(format string, args ...interface{}) {
	if logEnable(LevelDebug) {
		log.SetPrefix(LevelDebug.String())
		log.Output(3, fmt.Sprintf(format, args...))
	}
}

func Info(args ...interface{}) {
	if logEnable(LevelInfo) {
		log.SetPrefix(LevelInfo.String())
		log.Output(3, fmt.Sprintln(args...))
	}
}

func Infof(format string, args ...interface{}) {
	if logEnable(LevelInfo) {
		log.SetPrefix(LevelInfo.String())
		log.Output(3, fmt.Sprintf(format, args...))
	}
}

func Warn(args ...interface{}) {
	if logEnable(LevelWarn) {
		log.SetPrefix(LevelWarn.String())
		log.Output(3, fmt.Sprintln(args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if logEnable(LevelWarn) {
		log.SetPrefix(LevelWarn.String())
		log.Output(3, fmt.Sprintf(format, args...))
	}
}

func Error(args ...interface{}) {
	if logEnable(LevelWarn) {
		log.SetPrefix(LevelInfo.String())
		log.Output(3, fmt.Sprintln(args...))
	}
}

func Errorf(format string, args ...interface{}) {
	if logEnable(LevelError) {
		log.SetPrefix(LevelError.String())
		log.Output(3, fmt.Sprintf(format, args...))
	}
}
