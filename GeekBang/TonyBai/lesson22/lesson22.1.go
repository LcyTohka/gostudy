package main

import (
	"errors"
	"fmt"
)

func main() {
	//一字符串形势呈现信息
	//err := errors.New("your first demo error")
	//errWithCtx = fmt.Errorf("index %d is out of bounds", i)
	err := doSomething()
	if err != nil {
		// 不关心err变量底层错误值所携带的具体上下文信息
		// 执行简单错误处理逻辑并返回
		fmt.Println(err)
	}
}

func doSomething() error {
	return errors.New("some error occurred")
}
