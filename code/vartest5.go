package main 

import "unsafe"
const(
	a = "admin"
	b = len(a)
	c = unsafe.Sizeof(a)
)
func main() {
	/*常量 测试 */
	println(a,b,c)
}