package basics

import (
	"fmt"
	"math"
)

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)

	const (
		js = iota
		ts
		php
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tp
		pb
	)
}

func consts() {
	const file = "abc.txt" // 一般不用uppercase 定义
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b)) // 常量如果没有定义类型不需要强转
	fmt.Println(file, c)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 包内部变量
// := 只能用在函数内部
var aa = 3

var (
	bb = 4
	ss = "kkk"
)

func variableInitValues() {
	var a, b int = 6, 8
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableInitValuesTypeDeduction() {
	var a, b, c = 6, 8, true
	var s = "abc"
	fmt.Println(a, s, b, c)
}

func variableInitValuesTypeDeduction2() {
	a, b, c := 6, 8, true
	var s = "abc"
	fmt.Println(a, s, b, c)
}

//func main() {
//	fmt.Println("hello, there")
//	fmt.Println(runtime.GOARCH)
//
//	//variableInitValuesTypeDeduction2()
//	triangle()
//}
