package mysort

import (
	"fmt"
	"goSnippets/logger"
)

// 最优的单向冒泡排序
func bubbleSort(a []int) {
	for n := len(a); n > 1; {
		newn := 0 // 用来保证一定退出最外层循环
		for i := 1; i < n; i++ {
			// 符合func Less(i, j int) bool
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				newn = i // 用来跳过比较已排序的元素
			}
		}
		n = newn
	}
}

func init() {
	logger.DefaultLogger.Log()
	a := []int{4, 2, 3, 5, 1}
	bubbleSort(a)
	fmt.Println(a)
}
