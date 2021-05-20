package apiprac

import (
	"fmt"
	"goSnippets/logger"
	"unicode/utf8"
)

// byte count & rune count
func CharCount(s string) (bytes, runes int) {
	bytes = len(s)
	runes = utf8.RuneCountInString(s)
	return
}

func init() {
	logger.DefaultLogger.Log()
	const msg = "你好中国，hello Chinese"
	b, r := CharCount(msg)
	fmt.Printf("bytes: %v, runes: %v\n", b, r)
}
