package tests

import (
	"GoDemoProj/testing/funcs"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct{ a, b, sum int }{
		{1, 2, 3},
		//{1, 1, 3}, // wrong case
		{3, 5, 8},
		{0, 0, 0},
	}

	for _, test := range tests {
		if actual := funcs.AddTwoNumbers(test.a, test.b); actual != test.sum {
			t.Errorf("add(%d, %d), got %d, expected %d", test.a, test.b, actual, test.sum)
		}
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	a1, a2 := 1, 100
	sum := 101
	// do some setup

	b.Logf("do some setup")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := funcs.AddTwoNumbers(a1, a2)
		if actual != sum {
			b.Errorf("add(%d, %d), got %d, expected %d", a1, a2, actual, sum)
		}
	}
}
