package main

import "fmt"

func main() {
	s := make(chan string, 2)
	s <- "admin"
	s <- "Jean"
	close(s) //如果没有 close 它，我们将在这个循环中继续阻塞执行，等待接收第三个值
	for item := range s {
		fmt.Println(item)
	}
}
