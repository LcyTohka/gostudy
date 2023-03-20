package main

import "fmt"

func main() {

	b := make([]int, 6, 8)
	fmt.Println(b, len(b), cap(b))
	//arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//sl := arr[3:7:9]
	////len(sl)7-3
	////cap(sl)9-3
	//fmt.Println(sl, len(sl), cap(sl))
	//
	////切片扩容
	//var s []int
	//s = append(s, 11)
	//fmt.Println(len(s), cap(s)) //1 1
	//s = append(s, 12)
	//fmt.Println(len(s), cap(s)) //2 2
	//s = append(s, 13)
	//fmt.Println(len(s), cap(s)) //3 4
	//s = append(s, 14)
	//fmt.Println(len(s), cap(s)) //4 4
	//s = append(s, 15)
	//fmt.Println(len(s), cap(s)) //5 8

	u := [...]int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	s = append(s, 24)
	fmt.Println("after append 24, array:", u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 25)
	fmt.Println("after append 25, array:", u)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 26)
	fmt.Println("after append 26, array:", u)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
}
