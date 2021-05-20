package main

import (
	"fmt"
	"time"
)

func main() {
	//demoChannel()

	//bufferedChannel()
	channelClose()
}

func channelClose() {
	c := make(chan int, 3)
	go workerCloseChannel(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)

	close(c)
	time.Sleep(time.Millisecond * 2)
}

func workerCloseChannel(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %d\n", id, n)
	//}

	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func workerBufferedChannel(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %d\n", id, <-c)
	}
}

func bufferedChannel() {
	c := make(chan int, 3)
	go workerBufferedChannel(0, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4 // without worker to consume channel, fatal error: all goroutines are asleep - deadlock!

	time.Sleep(time.Millisecond * 2)
}

// chan<- int, send only channel
// <-chan int, receive only channel

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()

	return c
}

func worker(id int, c chan int) {
	for {
		//n := <-c
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func demoChannel() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond * 2)
}
