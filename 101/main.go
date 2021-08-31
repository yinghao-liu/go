package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex sync.RWMutex
var ch chan int = make(chan int) // 无缓冲

func readOP(tick int) {
	fmt.Printf("%v before rlock, %v s\n", tick, time.Now().Second())
	rwMutex.RLock()
	fmt.Printf("%v -----after rlock, %v s\n", tick, time.Now().Second())
	defer rwMutex.RUnlock()

	// work
	time.Sleep(5e9)

}

func writeOP(tick int) {
	fmt.Printf("%v before wlock, %v s\n", tick, time.Now().Second())
	rwMutex.Lock()
	fmt.Printf("%v -----after wlock, %v s\n", tick, time.Now().Second())
	defer rwMutex.Unlock()

	// work
	time.Sleep(5e9)

}

func main() {
	for i := 0; i < 10; i++ {
		if 6 == i {
			go writeOP(i)
		} else {
			go readOP(i)
		}
		time.Sleep(1e8)
	}
	for {
		time.Sleep(5e9)
	}
}
