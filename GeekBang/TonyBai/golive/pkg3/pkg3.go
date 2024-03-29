package pkg3

import (
	"fmt"
)

var (
	_ = constInitCheck()
	v = variableInit("v")
)

const (
	c = "c"
)

func constInitCheck() string {
	fmt.Println("pkg3: const c has been initialized")

	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg3: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg3: init func invoked")
}

func main() {
	// do nothing
}
