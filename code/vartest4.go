package main 

import "fmt"
func main() {
	/*常量声明测试 */
	const length int = 50
	const width int = 5
	var area int
	const a,b,c = 1,false,"str"   //多重赋值

	area = length*width
	fmt.Printf("面积为:%d",area)
	println()
	println(a,b,c)
}