package main

import "fmt"

func main() {
	messages := make(chan string, 2) // 这个channel只是接受两个字符串   所以 可以用来做缓存

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
