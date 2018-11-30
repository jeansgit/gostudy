package main

import "fmt"

type test struct {
	name string
	age  int
}

func main() {
	fmt.Println(test{"admin", 100})
	fmt.Println(test{name: "jean"})
	fmt.Println(&test{"haha", 300})

	x := test{"good", 100}
	fmt.Println(x.name)
}
