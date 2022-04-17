package utility

import (
	"fmt"
	"runtime"
)

// 调用堆栈
func CallStack() {
	pc := make([]uintptr, 8)
	num := runtime.Callers(2, pc)
	pc = pc[:num-2]
	frames := runtime.CallersFrames(pc)
	if nil == frames {
		return
	}
	fmt.Printf("-----------------CallStack-------------------\n")
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d,func:%s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
	fmt.Printf("-----------------CallStack-------------------\n")
}
