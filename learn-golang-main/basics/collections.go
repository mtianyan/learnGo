package basics

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//array()
	//slice()
	//sliceOperation()
	//sliceOps()
	//maps()
	//chars()
	//s1 := "hello"
	//s2 := "hello"
	//s3 := "hello!"
	//println(s1 == s2)
	//println(s2 == s3)
}

func chars() {
	s := "Yes你好啊"
	println(len(s)) // len = 12
	for _, b := range []byte(s) {
		fmt.Printf("%s : %X \n", string(b), b)
	}

	for i, ch := range s {
		fmt.Printf("index: %d, value: %X \n", i, ch)
	}

	println("Rune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	println()

	// rune 相当于Go 里面的char，支持utf8
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
		print()
	}
	println()
}

func maps() {
	map1 := map[string]string{
		"name":    "hi",
		"address": "world",
	}

	map2 := make(map[string]int) // empty map
	var map3 map[string]int      // nil
	fmt.Println("map1: ", map1)
	fmt.Println("map2: ", map2, map2 == nil)
	fmt.Println("map3: ", map3, map3 == nil)

	//map1:  map[address:world name:hi]
	//map2:  map[] false
	//map3:  map[] true

	// iterate map entries
	// hashmap ->  unordered
	// iteration order may be different for different runs
	for k, v := range map1 {
		println(k, v)
	}

	// get values
	println("get name from map: ", map1["name"])
	println("get name from map: ", map1["nam"]) // non exists entry, empty (zero value)

	//randomName := "randomName"
	if randomName, ok := map1["randomName"]; ok {
		println("randomName: ", randomName)
	} else {
		println("key does not exist")
	}

	someName, ok := map1["haha"]
	println("value, ok: ", someName, ok)

	// delete values
	delete(map1, "name")
	value, exist := map1["name"]
	println("after deletion: ", value, exist)
}

func sliceOps() {
	// 创建slice
	var s []int // default init value for slice is nil
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len = %d, cap = %d \n", len(s), cap(s))
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	fmt.Println(s1)

	s2 := make([]int, 16) // with len
	fmt.Printf("len = %d, cap = %d, slice2 = %v \n", len(s2), cap(s2), s2)

	s3 := make([]int, 16, 20) // with len and capacity
	fmt.Printf("len = %d, cap = %d, slice3 = %v \n", len(s3), cap(s3), s3)

	// copy slice
	copy(s2, s1)
	fmt.Println("s2=", s2)

	// delete elements from slice
	// delete element with index 3
	s2 = append(s2[:3], s2[4:]...)
	s2 = s2[1:]         // remove first
	s2 = s2[:len(s2)-1] // remove last
}

func sliceOperation() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := array[2:6]
	s1 := append(s, 100)
	s2 := append(s1, 101)
	s3 := append(s2, 102)
	fmt.Println(s, s1, s2, s3)
	fmt.Println(array)
	// [2 3 4 5] [2 3 4 5 100] [2 3 4 5 100 101] [2 3 4 5 100 101 102]
	// [0 1 2 3 4 5 100 101]

}

func slice() { // slice is a view of array
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7} // array type
	s := array[2:6]                           // s is slice type, index 6 is excluded
	fmt.Println(s)
	fmt.Println("array[:6]", array[:6])
	fmt.Println("array[2:]", array[2:])
	fmt.Println("array[:]", array[:])

	s = s[0:3] // re-slice
	// extend slice
	s1 := array[2:6]
	s2 := s1[3:5] // 5 is beyond s1's len,
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	println("cap s1:", cap(s1))
	println("len s1:", len(s1))
}

func array() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5} // when using :=, must init array
	arr3 := [...]int{2, 4, 6}
	var grid [2][3]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		println(arr3[i])
	}

	for i := range arr3 {
		println(arr3[i])
	}

	for i, v := range arr3 {
		println(i, v)
	}
}

func printArray(arr [3]int) {

}

func printArrayRef(arr *[5]int) {

}
