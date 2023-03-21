package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

func main() {

	type Empty struct{} // Empty是一个不包含任何字段的空结构体类型

	var s Empty
	println(unsafe.Sizeof(s)) // 0

	var b bytes.Buffer
	b.Write([]byte("Hello, Go"))
	fmt.Println(b.String()) // 输出：Hello, Go
}
