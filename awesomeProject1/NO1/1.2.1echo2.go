package main

import (
	"fmt"
	"os"
)

func main() {
	for temp1, arg := range os.Args {
		fmt.Println(temp1, arg)
	}
}
