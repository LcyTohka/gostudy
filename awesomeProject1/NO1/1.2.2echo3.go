package main

import (
	"fmt"
	_ "fmt"
	"os"
	_ "os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args)
}
