package main

import "fmt"
import "os"

func main() {
	cmd1 := os.Args
	cmd2 := os.Args[3:]
	cmd3 := os.Args[3]
	fmt.Println(cmd1)
	fmt.Println(cmd2)
	fmt.Println(cmd3)
}
