package main

import (
	"fmt"
	"io"
	"strings"
)

func NextFibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func NextFibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 实现Reader 接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()      // call next fib
	if next > 1000 { // end reading next
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}
