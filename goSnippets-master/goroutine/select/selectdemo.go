package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p1, p2 := producer(), producer()
	consumer := newConsumer()

	var values []int
	timer := time.After(10 * time.Second)
	ticker := time.Tick(time.Second)
	for {
		var activeConsumer chan<- int
		var activeValue int
		if len(values) > 0 {
			activeConsumer = consumer
			activeValue = values[0]
		}
		select {
		case n := <-p1:
			values = append(values, n)
		case n := <-p2:
			values = append(values, n)
		case activeConsumer <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-ticker:
			fmt.Println("queue len =", len(values))
		case <-timer:
			fmt.Println("deadline")
			return
		}
	}
}

func newConsumer() chan<- int {
	con := make(chan int)
	go consume(con)
	return con
}

func consume(c chan int) {
	for n := range c {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Consumer consume %d\n", n)
	}
}

func producer() <-chan int {
	pro := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			pro <- i
			i++
		}
	}()
	return pro
}
