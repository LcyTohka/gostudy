package main

import "fmt"

func add(a []int, i int, j int) []int {
	b := a[:i]
	c := []int{}
	fmt.Println(a[i:])
	for k := 1; k <= len(a[i:]); k++ {
		c = append(c, a[i+k-1])
	}
	fmt.Println(c)
	b = append(b, j)
	b = append(b, c...)
	return b
}

//func delete(i int, j int) []int {
//	return
//}

func main() {
	b := []int{1, 2, 3, 5}
	a := add(b, 3, 4)
	fmt.Println(a)

}
