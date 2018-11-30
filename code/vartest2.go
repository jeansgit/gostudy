package main 
var a,b int 
var (	//一般用于全局变量声明
	c int
	d bool
	e string
)

var f,g int = 1,2
var h,i = 123,"test"

func main() {
	/*变量声明测试*/
	j,k := 456,"jean"
	println(a,b,c,d,e,f,g,h,i,j,k)
}