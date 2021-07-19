package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println("producer------", i)
		ch <- i
	}
}

func customer(flag int, ch chan int) {
	for {
		data := <-ch
		fmt.Println(flag, "---customer---", data)
		time.Sleep(1e9)
	}
}

/*
	like in a real world, at a moment, large numbers of requests rush in,
	but work is hard and slow
*/
func main() {
	ch := make(chan int, 10)
	go producer(ch)
	time.Sleep(1e9)

	// create two customers
	go customer(1, ch)
	go customer(2, ch)
	go customer(3, ch)

	for {
		time.Sleep(10)
	}
}
