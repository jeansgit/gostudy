package main
import "fmt"
func main() {
	nums := []int{1,2,3,4}
	sum := 0
	//使用slice求和
	for num := range nums{
		sum += num
	}
	fmt.Printf("%d\n",sum)
	for i,num:=range nums{
		if(num==3){
			fmt.Printf("Index:%d\n",i)
		}
	}
	res := map[string]string{"a":"admin","b":"jean","c":"haha"}
	for k,v:=range res{
		fmt.Printf("%s-%s\n",k,v)
	}
	for i,char:=range "admin"{
		fmt.Println(i,char)
	}
}