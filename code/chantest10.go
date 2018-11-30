package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := "admin":
		fmt.Println("received message", msg)
	}
	msg := "hi"
	//msg2 := "adminjjj"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)

	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)

	}
}
