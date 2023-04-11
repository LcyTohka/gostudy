package main

func main() {

	n := <- ch      // 当ch被关闭后，n将被赋值为ch元素类型的零值
	m, ok := <-ch   // 当ch被关闭后，m将被赋值为ch元素类型的零值, ok值为false
	for v := range ch { // 当ch被关闭后，for range循环结束
		... ...
	}
}
