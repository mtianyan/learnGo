package mymath

import (
	"fmt"
	"goSnippets/logger"
)

const Len = 41

var cache [Len]int

func CachedFib(n int) int {
	if n <= 1 {
		return n
	}
	if cache[n] == 0 {
		cache[n] = CachedFib(n-1) + CachedFib(n-2)
	}
	return cache[n]
}

func init() {
	logger.DefaultLogger.Log()
	for i := 0; i < Len; i++ {
		fmt.Printf("%v: %v\n", i, CachedFib(i))
	}
}
