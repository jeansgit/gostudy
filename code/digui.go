package main
import "fmt"
func main() {
	for i:=0;i<20;i++{
		fmt.Printf("%d的阶乘是:%d \n",i,jiecheng(uint64(i)))	
	}
}
func jiecheng(n uint64)(result uint64) {
	if(n>0){
		result = n*jiecheng(n-1)
		return result
	}
	return 1
}