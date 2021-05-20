package queue

import "fmt"

func ExampleQueue_Poll() {
	q := Queue{1}
	q.Push(0)

	fmt.Println(q.Poll())
	// Output:
	// 1
}
