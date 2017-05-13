package main

import "fmt"
import "time"

// 使用select case 来选择channel 执行
func main() {
	c1 := make(chan string) //定义channel
	c2 := make(chan string)
	go func() { // goroutine中向channel中发送数据
		time.Sleep(time.Second * 4)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("in for loop")

		select {
		case msg1 := <-c1:
			fmt.Println("msg1 received", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2 received", msg2)
		}
	}
}
