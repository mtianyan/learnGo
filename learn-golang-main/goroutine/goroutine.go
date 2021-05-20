package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("hello from goroutine %d \n", i)
				a[i]++
				runtime.Gosched() // 交出goroutine 控制权
			}
		}(i) // 使用传入i而不是直接取i
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
