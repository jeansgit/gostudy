package main

import "fmt"
import "time"

func work(msg chan bool) {
	fmt.Print("working......A")
	time.Sleep(time.Second)
	fmt.Println("done")

	msg <- true
}
func main() {
	msg := make(chan bool, 1)
	go work(msg)
	<-msg
}
