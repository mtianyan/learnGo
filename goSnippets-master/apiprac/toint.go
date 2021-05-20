package apiprac

import (
	"fmt"
	"goSnippets/logger"
	"math"
)

// 安全的float64转成int32
func SafeFtoi(f float64) int32 {

	if math.MinInt32 <= f && f <= math.MaxInt32 {
		// why not int(f+0.5)?
		i, frac := math.Modf(f)
		if frac >= 0.5 {
			i++
		}
		return int32(i)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", f))
}

func init() {
	logger.DefaultLogger.Log()
	v := 32.2
	fmt.Printf("%f to %d\n", v, SafeFtoi(v))
	v = 2.147483648123e+09

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	SafeFtoi(v)
}
