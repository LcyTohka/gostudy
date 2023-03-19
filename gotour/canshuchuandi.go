package main

import "fmt"

type cb func(int) int

func main() {
	a, b := 1, 2
	testCallBack(&a, callback)
	testCallBack(&b, func(x int) int {
		fmt.Printf("2: %d\n", x)
		return x
	})
}

func testCallBack(x *int, f cb) {
	f(*x)
}

func callback(x int) int {
	fmt.Printf("1: %d\n", x)
	return x
}
