package main

import "fmt"

//方法的一般生命形式
//func (t *T或T) MethodName(参数列表) (返回值列表) {
//	// 方法体
//}
//
//type T struct{}
//
//func (t T) M(n int) {
//}
//
//func main() {
//	var t T
//	t.M(1) // 通过类型T的变量实例调用方法M
//
//	p := &T{}
//	p.M(2) // 通过类型*T的变量实例调用方法M
//}

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

//func main() {
//	var t T
//	t.a = 123
//	fmt.Println(T.Get(t))
//}

func main() {
	var t T
	f1 := (*T).Set                           // f1的类型，也是*T类型Set方法的类型：func (t *T, int)int
	f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
	f1(&t, 3)
	fmt.Println(f2(t)) // 3
}
