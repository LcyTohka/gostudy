package main

import "fmt"

var (
	aa = 11
	bb = 22
	cc = "aa"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
}
