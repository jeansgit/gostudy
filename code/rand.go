package main

import "fmt"

//import "time"
import "math/rand"

func main() {
	p := fmt.Println
	p(rand.Intn(100))
	p(rand.Intn(100))
	p(rand.Intn(100))
	p(rand.Intn(100))
	p(rand.Float64())

}
