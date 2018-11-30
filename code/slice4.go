package main

import "fmt"
func main() {
	var numbers []int	//创建空切片
	printSlice(numbers)	//len=0  cap=0   slice=[]

	numbers = append(numbers,0) //向切片追加一个元素0 
	printSlice(numbers) //len=1  cap=1   slice=[0]

	numbers = append(numbers,2,3,4,5) //向切片增加多个元素
	printSlice(numbers) //len=5  cap=6   slice=[0 2 3 4 5]

	//创建一个长度不变，容量为numbers两倍的切片
	number1:=make([]int,len(numbers),(cap(numbers))*2) 

	copy(number1,numbers)	//拷贝numbers的内容到number1
	printSlice(number1)		//len=5  cap=12   slice=[0 2 3 4 5]
}

func printSlice(x []int) {
	fmt.Printf("len=%d  cap=%d   slice=%v\n",len(x),cap(x),x)
}