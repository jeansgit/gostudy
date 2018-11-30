package main

import "fmt"
var a int = 20

func main() {

	var a int = 10
	var b int = 20
	var c int = 0
	fmt.Printf("%d\n",a)
	c = add(a,b)
	fmt.Printf("%d\n",c)
}

func add(a,b int) int{
	fmt.Printf("%d\n",a)
	fmt.Printf("%d\n",b)
	return a+b
}