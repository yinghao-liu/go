package sub

import (
	"fmt"
	"runtime"
)

func Sub() {
	pc := make([]uintptr, 12)
	num := runtime.Callers(2, pc)
	fmt.Printf("num is %d\n", num)
	fmt.Printf("ps is %+v\n", pc)
	pc = pc[:num-2]
	frames := runtime.CallersFrames(pc)
	if nil == frames {
		return
	}
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d,func:%s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}

}
