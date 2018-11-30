package main

import "fmt"
import "time"

func main() {
	a1 := make(chan string, 1)
	a2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		a1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		a2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg := <-a1:
			fmt.Println(msg)
		case msg := <-a2:
			fmt.Println(msg)
		}
	}
}
