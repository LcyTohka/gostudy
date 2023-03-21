package main

import "fmt"

//func main() {
//
//	var sum int
//	var sl = []int{1, 2, 3, 4, 5, 6}
//	for i := 0; i < len(sl); i++ {
//		if sl[i]%2 == 0 {
//			// 忽略切片中值为偶数的元素
//			continue
//		}
//		sum += sl[i]
//	}
//	println(sum) // 9
//}

//func main() {
//	var sum int
//	var sl = []int{1, 2, 3, 4, 5, 6}
//
//loop:
//	for i := 0; i < len(sl); i++ {
//		if sl[i]%2 == 0 {
//			// 忽略切片中值为偶数的元素
//			continue loop
//		}
//		sum += sl[i]
//	}
//	println(sum) // 9
//}

func main() {
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == 13 {
				fmt.Printf("found 13 at [%d, %d]\n", i, j)
				break
			}
		}
	}
}
