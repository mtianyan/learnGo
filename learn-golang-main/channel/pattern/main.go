package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//tryMsgGen(msgGen())
	//tryNonBlockingWait(msgGen(), msgGen())
	tryMsgGenWithExist()
}

func tryMsgGenWithExist() {
	done := make(chan struct{})
	m1 := msgGenWithExit(done)
	for i := 0; i < 10; i++ {
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println("service1 get message: ", m)
		} else {
			fmt.Println("timeout!")
		}
	}

	done <- struct{}{}
	time.Sleep(time.Second)
	<-done // use done channel as bi-direction channel
}

func msgGenWithExit(done chan struct{}) chan string {
	c := make(chan string)

	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("Message %d", i)
			case <-done:
				fmt.Println("cleaning up now...")
				time.Sleep(time.Second)
				fmt.Println("cleaning up done!")
				done <- struct{}{}
				return
			}
			i++
		}
	}()

	return c
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func tryNonBlockingWait(m1, m2 chan string) {
	for {
		fmt.Println("service1:", <-m1)
		if m, ok := nonBlockingWait(m2); ok {
			fmt.Println("service2:", m)
		} else {
			fmt.Println("no message from service2")
		}
	}
}

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func fanInMultiChansFunc(chans ...chan string) chan string {
	c := make(chan string)
	for _, channel := range chans {
		// iterate过程中channel变量会被不断重新赋值
		// 利用函数传参的值拷贝
		go func(ch chan string) {
			c <- <-ch
		}(channel)
	}
	return c
}

func fanInMultiChans(chans ...chan string) chan string {
	c := make(chan string)
	for _, channel := range chans { // iterate过程中channel变量会被不断重新赋值
		// wrong!!
		// 外层for循环结束的时候，变量channel会停留在最后一个遍历到的channel
		// 此时所有的goroutine 都会使用这个channel 变量来执行
		// solution: channel copy
		channelCopy := channel
		go func() {
			c <- <-channelCopy
		}()
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case msg := <-c1:
				c <- msg
			case msg := <-c2:
				c <- msg
			}
		}
	}()

	return c
}

func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		c <- <-c1
	}()

	go func() {
		c <- <-c2
	}()

	return c
}

func tryMsgGen(c chan string) {
	for {
		fmt.Println("Got message: ", <-c)
	}
}

func msgGen() chan string {
	c := make(chan string)

	go func() {
		i := 0
		for {
			timeToSleep := rand.Intn(2000)
			time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
			c <- fmt.Sprintf("Message %d", i)
			i++
		}
	}()

	return c
}
