package main

import "fmt"

func main() {
    const (
        a=iota      //0
        b           //1
        c           //2
        d = "admin" // iota+1 = 3 
        e           // iota+1 = 4
        f = 100     // iota+1 = 5
        g           // iota+1 = 6
        h = iota    // iota+1 = 7
        i           // iota+1 = 8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}