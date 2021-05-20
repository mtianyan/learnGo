package apiprac

import (
	"fmt"
	"goSnippets/logger"
)

// 迭代Unicode字符串
func IterUnicodeStr(s string) {
	for _, char := range s {
		fmt.Printf("%c", char)
	}
	fmt.Println()
}

func init() {
	logger.DefaultLogger.Log()
	IterUnicodeStr("你好中国")
}
