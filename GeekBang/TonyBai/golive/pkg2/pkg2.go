package pkg2

import (
	"fmt"

	_ "github.com/LcyTohka/gostudy/tree/main/GeekBang/TonyBai/golive/pkg3"
)

var (
	_  = constInitCheck()
	v2 = variableInit("v2")
)

const (
	c2 = "c2"
)

func constInitCheck() string {
	if c2 != "" {
		fmt.Println("pkg2: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg2: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg2: first init func invoked")
}

func main() {
	// do nothing
}
