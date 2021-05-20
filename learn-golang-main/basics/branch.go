package basics

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
)

//func main() {
//	//readFile2()
//	//apply(plus, 1, 3)
//	//apply(func(a int, b int) int {
//	//	return a + b
//	//}, 100, 1)
//
//	a, b := 3, 4
//	swap_ref(&a, &b)
//	a, b = swap_val(a, b)
//	println(a, b)
//
//}

func swap_ref(a, b *int) {
	*a, *b = *b, *a
}

func swap_val(a, b int) (int, int) {
	return b, a
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func plus(a, b int) int {
	return a + b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)", opName, a, b)
	return op(a, b)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("invalid input") // 中断程序并且报错
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func eval(a, b int, op string) int {
	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		panic("Unsupported operation: " + op)
	}
	return res
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div1(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func readFile() {
	const filename = "go.mod"
	content, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Printf("%s \n", content)
	} else {
		fmt.Println(err)
	}
}

func readFile2() {
	const filename = "go.mod"
	// define like a for loop in java -> varible cannot be access outside if block
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", content)
	}
}
