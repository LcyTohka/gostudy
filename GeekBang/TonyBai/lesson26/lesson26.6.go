package main

import (
	"fmt"
	"reflect"
)

type W struct{}

func (T) M1()  {}
func (*T) M2() {}

type T11 W

func dumpMethodSet2(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
func main() {
	var t T
	var pt *T
	var t1 T11
	var pt1 *T11

	dumpMethodSet2(t)
	dumpMethodSet2(t1)

	dumpMethodSet2(pt)
	dumpMethodSet2(pt1)
}
