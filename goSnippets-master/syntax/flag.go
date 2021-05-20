package syntax

import (
	"fmt"
	"goSnippets/logger"
)

// 表示资源的使用状态
const (
	Open = 1 << iota
	Close
	Pending
)

func init() {
	logger.DefaultLogger.Log()
	fmt.Println("Open:", Open)
	fmt.Println("Close:", Close)
	fmt.Println("Pending:", Pending)
}
