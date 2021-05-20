package syntax

import (
	"fmt"
	"goSnippets/logger"
)

// 反引用 空指针
func NilPointer() {
	var p *int = nil
	fmt.Println(*p)
}

func init() {
	logger.DefaultLogger.Log()
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("*(nil pointer)")
		}
	}()
	NilPointer()
}
