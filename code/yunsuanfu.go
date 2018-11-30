package main

import "fmt"

func main() {
 var a int = 21
 var b int = 10
 var c int
 c = a + b
 fmt.Printf("a+b :%d\n",c)
 c = a - b
 fmt.Printf("a-b :%d\n",c)
 c = a * b
 fmt.Printf("a*b :%d\n",c)
 c= a / b
 fmt.Printf("a/b :%d\n",c)
 c = a % b
 fmt.Printf("a%%b :%d\n",c)
 a++
 fmt.Printf("a++ :%d\n",a)
 b--
 fmt.Printf ("b-- :%d\n",b)

 }