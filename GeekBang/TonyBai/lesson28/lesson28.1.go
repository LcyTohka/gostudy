package main

import "fmt"

//接口类型是由 type 和 interface 关键字定义的一组方法集合

func main() {
	var i interface{} // ok
	//type i = interface{}
	i = 15
	i = "hello, golang" // ok
	type T struct{}
	var t T
	i = t  // ok
	i = &t // ok
	fmt.Println(i)
}
