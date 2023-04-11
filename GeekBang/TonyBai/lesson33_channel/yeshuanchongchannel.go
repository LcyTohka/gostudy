package main

func main() {

	ch2 := make(chan int, 1)
	n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起

	ch3 := make(chan int, 1)
	ch3 <- 17 // 向ch3发送一个整型数17
	ch3 <- 27 // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起
}
