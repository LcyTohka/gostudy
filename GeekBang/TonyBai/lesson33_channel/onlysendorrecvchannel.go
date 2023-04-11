package main

func main() {

	ch1 := make(chan<- int, 1) // 只发送channel类型
	ch2 := make(<-chan int, 1) // 只接收channel类型

	<-ch1 // invalid operation: <-ch1 (receive from send-only type chan<- int)
	ch1 <- 17
	ch2 <- 13 // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
	<-ch2
}
