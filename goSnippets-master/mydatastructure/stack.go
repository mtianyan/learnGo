package mydatastructure

import (
	"fmt"
	"reflect"
)

func NewStack() *Stack {
	return &Stack{data: make([]Element, 16), top: 0}
}

type Element interface{}
type Stack struct {
	data []Element
	top  int
}

func (s *Stack) Push(value Element) {
	if !s.Empty() && reflect.TypeOf(value) != reflect.TypeOf(s.LastElement()) {
		panic(fmt.Errorf("push a wrong type value(%T) to stack(%T)",
			value, s.LastElement()))
	}
	if s.top == len(s.data) {
		s.data = append(s.data, make([]Element, len(s.data))...)
	}
	s.data[s.top] = value
	s.top++
}

func (s *Stack) Pop() Element {
	if s.Empty() {
		panic("pop on empty stack")
	}
	s.top--
	return s.data[s.top]
}

func (s *Stack) LastElement() Element {
	return s.data[s.top-1]
}

func (s *Stack) Empty() bool {
	return s.top == 0
}
