package main
import "fmt"

//通过 Go 语言的递归函数实现斐波那契数列
func fibo(n int) int {
	if(n<2){
		return n
	}
	return fibo(n-1)+fibo(n-2)
}

func main() {
	for i:=0;i<20;i++{
		fmt.Printf("%d ",fibo(i))
	}
}