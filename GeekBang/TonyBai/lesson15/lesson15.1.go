package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//var arr2 = [6]int {
	//	11, 12, 13, 14, 15, 16,
	//} // [11 12 13 14 15 16]

	var arr3 = [...]int{
		21, 22, 23,
	} // [21 22 23]
	fmt.Printf("%T\n", arr3) // [3]int

	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度：", len(arr))           // 6
	fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48
}
