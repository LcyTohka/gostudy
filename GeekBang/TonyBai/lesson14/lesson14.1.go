package main

import "fmt"

const (
	_, _ = iota, iota
	Apple, Banana
	Strawberry, Grape
	Pear, Watermelon
)

func main() {
	fmt.Println(Apple, Banana, Strawberry, Grape, Pear, Watermelon)
}
