package main

import "fmt"

func main() {
	var s = "中国人"
	for i, v := range s {
		fmt.Println(i, v)
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
