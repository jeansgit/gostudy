package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int 
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)

	/*  & 和 * 运算符实例 */
	ptr = &a
	fmt.Printf("%d\n",a)
	fmt.Printf("%d\n",*ptr)
}