package main

import "fmt"

func main() {
	messages := make(chan string, 3)
	messages <- "admin"
	messages <- "test"
	fmt.Println(messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
