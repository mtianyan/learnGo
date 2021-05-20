package mymath

import (
	"fmt"
	"goSnippets/logger"
	"math/big"
)

// 阶乘算法
func fac(n int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		res.Mul(res, big.NewInt(i))
	}
	return res
}

func init() {
	logger.DefaultLogger.Log()
	for i := 0; i <= 30; i++ {
		fmt.Printf("%d: %v\n", i, fac(int64(i)))
	}
}
