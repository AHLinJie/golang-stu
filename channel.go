jackage main

import (
	"fmt"
	"time"
)

type Sender chan<- int

type Receiver <-chan int

func abc(x, int) {
	return x
}

func main() {

	// 简单双向通道

	ch1 := make(chan string, 5)
	ch1 <- "value1"
	ch1 <- "value2"
	ch1 <- "value3"

	value := <-ch1
	value2 := <-ch1
	close(ch1)
	value3 := <-ch1

	fmt.Printf(value)
	fmt.Printf(value2)
	fmt.Printf(value3)

	// 单向通道
	var myChannel = make(chan int, (0))
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)

	v := abc(1)
}
