package logfmt

import (
	"fmt"
	"runtime"
	"strings"
)

type Log struct {
}

func (l Log) Fatal(format string, a ...interface{}) {
	fmt.Printf(l.formatPrefix()+format, a...)
}
func (l Log) Error(format string, a ...interface{}) {
	fmt.Printf(l.formatPrefix()+format, a...)
}
func (l Log) Warning(format string, a ...interface{}) {
	fmt.Printf(l.formatPrefix()+format, a...)
}
func (l Log) Info(format string, a ...interface{}) {
	fmt.Printf(l.formatPrefix()+format, a...)
}
func (l Log) Debug(format string, a ...interface{}) {
	fmt.Printf(l.formatPrefix()+format, a...)
}

func (l Log) formatPrefix() (prefix string) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fc := runtime.FuncForPC(pc)
		funName := strings.Split(fc.Name(), ".")
		return fmt.Sprintf("%s:%d  %s : ", file, line, funName[len(funName)-1])
	}
	return
}
