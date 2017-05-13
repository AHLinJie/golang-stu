package main

import "fmt"

// 关闭channel后其实还是可以取出里面的value
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
