package main

import "fmt"
func main() {
	var a int = 10		//声明实际变量
	var ip *int			//声明指针变量
	
	fmt.Printf("指针ip的值为:%x\n",ip)
	ip = &a			
	fmt.Printf("a变量的地址为:%x\n",&a)
	fmt.Printf("指针变量ip的地址为:%x\n",ip)
	fmt.Printf("指针变量的值为:%d\n",*ip)
}