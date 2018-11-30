package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	m["admin"] = 100
	m["test"] = 200
	fmt.Println(m)

	fmt.Println(len(m)) //计算长度

	Adelete(m, "admin") //delete移出元素
	fmt.Println(m)
}
