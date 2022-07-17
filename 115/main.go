package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}
func main() {
	ticker()
}

func after() {

	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}

}

// ticker
func ticker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
func Probe() {
	ticker := time.NewTicker(time.Second * time.Duration(ProbeSpan))

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("rtsp Probe - Current time: ", t)
				r.ProbeTask()
			}
		}
	}()
}
