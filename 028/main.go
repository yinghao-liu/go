package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	defer close(ch)
	for i := 0; i < 20; i++ {
		ch <- i
		fmt.Println("producer------", i)
	}
}

func customer1(flag int, ch chan int) {
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println(flag, "---customer1--- channel may be closed")
			break
		}
		fmt.Println(flag, "---customer1---", data)
		time.Sleep(1e9)
	}
}

// for-range and channel
func customer2(flag int, ch chan int) {
	for i := range ch {
		fmt.Println(flag, "---customer2---", i)
		time.Sleep(1e9)
	}
	fmt.Println(flag, "---customer2--- channel may be closed")
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
	go customer1(1, ch)
	go customer1(2, ch)
	go customer2(3, ch)
	go customer2(4, ch)

	time.Sleep(10e9)

	time.Sleep(10e9)

}
