package main

func main() {
	//对无缓冲 channel 类型的发送与接收操作，一定要放在两个不同的 Goroutine 中进行，否则会导致 deadlock
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	m := <-ch1
	println(m)
}
