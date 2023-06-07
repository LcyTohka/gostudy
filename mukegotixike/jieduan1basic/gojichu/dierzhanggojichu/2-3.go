package main

import (
	"fmt"
	"math"
)

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
	)
	fmt.Println(cpp, python, golang)
}

func main() {
	enums()
}
