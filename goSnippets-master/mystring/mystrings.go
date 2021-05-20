package mystring

import "fmt"

// 修改字符串中的字符
func Set(s *string, index int, char byte) {
	temp := []byte(*s)
	temp[index] = char
	*s = string(temp)
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func init() {
	s := "Yes!你好中国"
	for i, ch := range s { // i 不连续，按byte计数
		fmt.Println(i, ch)
	}
	for i, ch := range []rune(s) { // i 连续，按rune计数
		fmt.Println(i, ch)
	}
}
