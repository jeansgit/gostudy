package main

import "fmt"
type Book struct{
	title string
	author string
	subject string
	bookid int
}
func main() {
	//创建一个结构体
	fmt.Println(Book{"test","admin","jean",10})

	//使用key-value格式赋值
	fmt.Println(Book{title:"test",author:"admin",subject:"jean",bookid:10})

	//忽略的字段为空，为0 
	fmt.Println(Book{title:"test",author:"admin"})	
}