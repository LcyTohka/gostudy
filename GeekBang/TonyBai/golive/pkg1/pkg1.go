package pkg1

import (
	"fmt"

	_ "github.com/LcyTohka/gostudy/tree/main/GeekBang/TonyBai/golive/pkg3"
)

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
)

const (
	c1 = "c1"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("pkg1: const c1 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg1: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg1: first init func invoked")
}

func main() {
	// do nothing
}
