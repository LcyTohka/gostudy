package main

import "fmt"

const Pi float64 = 3.14159265358979323846 // 单行常量声明

// 以const代码块形式声明常量
const (
	size    int64 = 4096
	i, j, s       = 13, 14, "bar" // 单行声明多个常量
)

type myInt int

const n = 13

func main() {
	var a myInt = 5
	fmt.Println(a + n) // 输出：18
}
