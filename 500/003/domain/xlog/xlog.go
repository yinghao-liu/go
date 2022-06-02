package xlog

var Log logger

type logger interface {
	Fatal(format string, a ...interface{})
	Error(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Info(format string, a ...interface{})
	Debug(format string, a ...interface{})
}

func Fatal(format string, a ...interface{}) {
	Log.Fatal(format, a...)
}
func Error(format string, a ...interface{}) {
	Log.Error(format, a...)
}
func Warning(format string, a ...interface{}) {
	Log.Warning(format, a...)
}
func Info(format string, a ...interface{}) {
	Log.Info(format, a...)
}
func Debug(format string, a ...interface{}) {
	Log.Debug(format, a...)
}
