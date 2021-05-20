package myslice

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func init() {
	a := []int{1, 2}
	b := []int{3, 4}

	// 拼接
	fmt.Println(append(a, b...))

	// 尾部扩容
	fmt.Println(cap(a), cap(append(a, make([]int, 10)...)))

	// 指定位置扩容
	i := 1
	fmt.Println(cap(a),
		cap(append(a[:i], append(make([]int, 10), a[i:]...)...)),
		append(a[:i], append(make([]int, 10), a[i:]...)...))

	// 指定位置删除
	i = 1
	fmt.Println(append(a[:1], a[i+1:]...))

	// 指定范围删除[i,j)
	i, j := 1, 2
	fmt.Println(append(a[:i], a[j:]...))

	// Pop
	fmt.Println(a[len(a)-1], a[:len(a)-1])

	// 插入单个元素
	x := 3
	fmt.Println(append(a[:0], append([]int{x}, a[0:]...)...))

	// 插入多个元素
	fmt.Println(append(a[:0], append(b, a[0:]...)...))
}

func FindDigits(filename string) []byte {
	file, _ := ioutil.ReadFile(filename)
	return regexp.MustCompile(`\d+`).Find(file) // 返回的切片引用了整个文件
}

func LowMemFindDigits(filename string) []byte {
	file, _ := ioutil.ReadFile(filename)
	digits := regexp.MustCompile(`\d+`).Find(file)
	cache := make([]byte, len(digits))
	copy(cache, digits) // 仅引用有效数据
	return cache
}
