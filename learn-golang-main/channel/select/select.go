package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//unsafeChannelExampleSolution()
	useTimer()
}

func useTimer() {
	c1 := generator()
	c2 := generator()

	var workerChan = createWorker(0)
	var values []int
	timer := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	timeout := time.After(800 * time.Millisecond)
	for {
		var activeWorkerChan chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorkerChan = workerChan
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorkerChan <- activeValue:
			values = values[1:]
		case <-timeout:
			println("timeout!")
		case <-tick:
			println("queue length is: ", len(values))
		case <-timer:
			println("bye~~")
			return
		}
	}
}

func unsafeChannelExampleSolution() {
	c1 := generator()
	c2 := generator()

	var workerChan = createWorker(0)
	var values []int
	for {
		var activeWorkerChan chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorkerChan = workerChan
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorkerChan <- activeValue:
			values = values[1:]
		}
	}
}

func unsafeChannelExample() {
	// in this case, since worker sleep 1 second after each run
	// during worker sleep c1 and c2 keep getting values from generator
	// in for loop, select block will receive values from c1 and c2 constantly
	// but those values won't be send to worker's channel so they will be "flushed"
	// solution: create a temp queue to hold unprocessed values and pull from queue
	c1 := generator()
	c2 := generator()

	var workerChan = createWorker(0)
	n := 0
	hasValue := false
	for {
		var activeWorkerChan chan<- int
		if hasValue {
			activeWorkerChan = workerChan
		}

		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorkerChan <- n:
			// this branch won't be selected
			// during worker sleep
			// this branch will only be selected
			// when worker is executing "n := range c" line
			hasValue = false
		}
	}
}

func worker(id int, c chan int) {
	for n := range c { // channel will be received only when this line got executed
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}
