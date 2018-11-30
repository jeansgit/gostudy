package main

import "fmt"

func main() {
    const (
        i = 1<<iota     //1<<0
        j = 3<<iota     //3<<1
        k               //3<<2
        l               //3<<3
    )
    fmt.Println(i)
    fmt.Println(j)
    fmt.Println(k)
    fmt.Println(l)
}