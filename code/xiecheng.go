package main

import "fmt"

func f(input string) {
	for i := 0; i < 3; i++ {
		fmt.Println(input, ":", i)
	}
}

func main() {
	f("xiecheng")

	go f("xiec")

	go func(msg string) {
		fmt.Println(msg)
	}("test")

	var input2 string
	fmt.Scanln(&input2)
	fmt.Println("done")
}
