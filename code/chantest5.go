package main

import "fmt"

//只允许发送数据的通道
func sendmsg(sender chan<- string, msg string) {
	sender <- msg
}

//receiver允许接收数据
//send允许发送数据

func receivemsg(receiver <-chan string, send chan<- string) {
	msg := <-receiver
	send <- msg
}
func main() {
	sender := make(chan string, 1)
	send := make(chan string, 1)
	sendmsg(sender, "admin")
	receivemsg(sender, send)A
	fmt.Println(<-send)
	fmt.Println(send)
}
