package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	var res [3][5]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			res[i][j] = i + j
		}
	}
	fmt.Println(res)
}
