package main

import "fmt"

//类型定义

type T1 int
type T2 T1
type T3 string

//type (
//	T1 int
//	T2 T1
//	T3 string
//)

func main() {
	//var n1 T1
	//var n2 T2 = 5
	//n1 = T1(n2) // ok
	//
	//var s T3 = "hello"
	//n1 = T1(s) // 错误：cannot convert s (type T3) to type T1

	type T = string

	var s string = "hello"
	var t T = s           // ok
	fmt.Printf("%T\n", t) // string
}
