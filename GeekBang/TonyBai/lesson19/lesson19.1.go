package main

import "fmt"

func main() {

	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)

	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum += i + j + k
		println(sum)
	}

	var sl = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("sl[%d] = %d\n", i, sl[i])
	}

	for i, v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}

	for i := range sl {
		fmt.Printf("sl[%d] = %d\n", i)
	}
	for _, v := range sl {
		fmt.Printf("sl[%d] = %d\n", v)
	}

	var s = "中国人"
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}

	var m = map[string]int{
		"Rob":  67,
		"Russ": 39,
		"John": 29,
	}

	for k, v := range m {
		println(k, v)
	}
}
