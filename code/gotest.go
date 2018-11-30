package main

import "fmt"
import "math"

const s string = "admintest"

func main() {
	fmt.Println(s)
	const n = 1000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
