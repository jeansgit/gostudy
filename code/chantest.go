package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "adminjean" }()

	msg := <-messages
	fmt.Println(msg)
}
