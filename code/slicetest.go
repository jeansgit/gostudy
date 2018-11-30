package main

import "fmt"

func main() {
	s := make([]string, 5)
	fmt.Println(s)

	slicetest := make([][]int, 3)
	for i := 0; i < 3; i++ {
		len := i + 1
		slicetest[i] = make([]int, len)
		for j := 0; j < len; j++ {
			slicetest[i][j] = i + j
		}
	}
	fmt.Println(slicetest)
}
