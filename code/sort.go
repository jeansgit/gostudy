package main

import "fmt"
import "sort"

func main() {
	strs := []string{"a", "c", "d", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{1, 32, 7, 4, 15}
	sort.Ints(ints)
	fmt.Println(ints)
	//检查是否已经排好序
	issort := sort.IntsAreSorted(ints)
	fmt.Println(issort)

}
