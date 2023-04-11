package main

func main() {
	//默认值为nil
	//var ch chan int

	ch1 := make(chan int) //无缓冲channel
	//ch2 := make(chan int, 5) //有缓冲channel

	ch1 <- 13  //将整型字面值13发送到无缓冲channel类型变量ch1中
	n := <-ch1 //从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
	println(n)
	//ch2 <- 17 //将整型字面值17发送到带缓冲channel类型变量ch2中
	//m := <- ch2//从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中
}
