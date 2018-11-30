package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums)
	total := 0
	for _, i := range nums {
		total = total + i
		//fmt.Println(total, i)
	}
	fmt.Println(total)
}
func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)

	numbers := []int{1, 2, 3, 4, 5, 6}
	sum(numbers...)
}
