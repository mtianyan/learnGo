package main

import (
	"GoDemoProj/errorhandling/fib"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//tryDefer()
	//writeFibToFile("./errorhandling/fib/fib.txt")
	//tryDefer2()
	writeFibToFile2("./errorhandling/fib/fib.txt")
}

func writeFibToFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//fmt.Println("File already exists!", err.Error())
		//return
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func writeFibToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func tryDefer() {
	defer println(1)
	defer println(2)
	println(3)
	panic("!!")
	// output above is 3, 2, 1
	// defered statement will be executed right before
	// function exit or before throwing out error
	// defer execution order is reversed
}

func tryDefer2() {
	for i := 0; i < 10; i++ {
		defer println(i)
		if i == 7 {
			break
		}
	}
}
