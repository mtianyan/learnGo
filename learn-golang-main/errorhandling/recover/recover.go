package main

import (
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() { // recover 只能在defer时调用
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is an error"))
	//a, b := 5, 0
	//fmt.Println(a / b) // runtime error: integer divide by zero

	//panic: 123 [recovered]
	//panic: 123
	panic(123)
}
