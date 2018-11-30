package main

import "fmt"
func main() {
	numbers:=[]int{0,1,2,3,4,5,6,7,8,9,10}
	printSlice(numbers)
	fmt.Printf("%d\n",numbers[1:3])
	fmt.Printf("%d\n",numbers[2:7])
	fmt.Printf("%d\n",numbers[:3])
	fmt.Printf("%d\n",numbers[4:])
	number1:=make([]int,0,5)
	number2:=numbers[:3]
	printSlice(number1)
	printSlice(number2)
	number3:=numbers[2:5]
	printSlice(number3)
	number4:=numbers[3:8]
	printSlice(number4)
}

func printSlice(x []int) {
	fmt.Printf("len=%d  cap=%d   slice=%v\n",len(x),cap(x),x)
}