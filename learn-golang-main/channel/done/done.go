package main

import (
	"fmt"
	"sync"
)

func main() {
	//demoChannel()
	demoWaitGroup()
}

type workerChannel struct {
	c    chan int
	done chan bool
}

func createWorker(id int) workerChannel {
	wc := workerChannel{
		c:    make(chan int),
		done: make(chan bool),
	}
	go worker(id, wc.c, wc.done)
	return wc
}

func worker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		//go channel 的send 是阻塞式的，这里会造成deadlock
		//done <- true
		//solution1:
		go func() {
			done <- true
		}()
	}
}

func demoChannel() {
	var channels [10]workerChannel
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	// create 20 jobs, 2 jobs per worker
	for i, channel := range channels {
		channel.c <- 'a' + i
	}

	for i, channel := range channels {
		channel.c <- 'A' + i
	}

	for _, wc := range channels {
		<-wc.done
		<-wc.done // 如果只有一个job，不会产生deadlock
	}
}

// -------------------------------------------------------------
// use wait group
type workerChanWg struct {
	c chan int
	//wg *sync.WaitGroup
	done func() // abstract as a function
}

func createWorkerWG(id int, wg *sync.WaitGroup) workerChanWg {
	wc := workerChanWg{
		c: make(chan int),
		//wg: wg,
		done: func() {
			wg.Done()
		},
	}
	go workerWG(id, wc)
	return wc
}

func workerWG(id int, wc workerChanWg) {
	for n := range wc.c {
		fmt.Printf("worker %d received %c\n", id, n)
		//done <- true // go channel 的send 是阻塞式的，这里会造成deadlock
		wc.done()
	}
}

func demoWaitGroup() {
	var waitGroup sync.WaitGroup // init??
	var workerChans [10]workerChanWg

	for i := 0; i < 10; i++ {
		workerChans[i] = createWorkerWG(i, &waitGroup)
	}

	waitGroup.Add(20)
	// create 20 jobs, 2 jobs per worker
	for i, channel := range workerChans {
		channel.c <- 'a' + i
	}

	for i, channel := range workerChans {
		channel.c <- 'A' + i
	}
	waitGroup.Wait()
}
