package main

import (
	"fmt"
)

func main() {

	m := make(map[int]string)
	m[1] = "value1"
	m[2] = "value2"
	m[3] = "value3"
	fmt.Println(m)

	v, ok := m[4]
	if !ok {
		fmt.Println("NO")
	} else {
		fmt.Println(v)
	}

	delete(m, 3)
	fmt.Println(m)

	for i, j := range m {
		fmt.Printf("[%d.%s]", i, j)
	}
}
