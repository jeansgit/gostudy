package main

import "fmt"

func bibao() func() int {
	i := 0
	return func() int {
		i = i + 1
		return i
	}
}
func main() {
	test := bibao()
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())

	test1 := bibao()
	fmt.Println(test1())
}
