package main

import "fmt"
var d int

func main() {
	var a,b,c int
	a = 10
	b = 20
	c = a + b
	d = a + b
	fmt.Printf("a+b(c为局部变量): %d\n",c)
	fmt.Printf("a+b(d为全局变量): %d\n",d)
}
