package mymath

import (
	"fmt"
	"goSnippets/logger"
	"math"
)

func GoldenFib(i int) int {
	j := (1 + math.Sqrt(5)) / 2
	return int(math.Floor(math.Pow(j, float64(i))/math.Sqrt(5) + 1/2))
}

func init() {
	logger.DefaultLogger.Log()
	for i := 0; i <= 30; i++ {
		fmt.Printf("%v: %v\n", i, GoldenFib(i))
	}
}
