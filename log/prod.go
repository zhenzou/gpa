// +build prod

package log

func SetLevel(level level) {}

func Debug(args ...interface{}) {}

func Debugf(format string, args ...interface{}) {}

func Info(args ...interface{}) {}

func Infof(format string, args ...interface{}) {}

func Warn(args ...interface{}) {}

func Warnf(format string, args ...interface{}) {}

func Error(args ...interface{}) {}

func Errorf(format string, args ...interface{}) {}
