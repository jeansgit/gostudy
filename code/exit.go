package main

import "fmt"
import "os"

func main() {
	var char = []string{"`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "=", "_",
		"+", ",", ".", "/", ";", "\\", "\"", "<", ">", "?", "|", ":", "[", "]", "{", "}",
		"!@", "!#", "@#", "@!", "#@", "#!", "!@#", "!#@", "#@!", "#!@", "@!#", "@#!", "!@#$",
		"!@#$%", "!@#$%^", "!@#$%^&", "!@#$%^&*", "!@#$%^&*(", "!@#$%^&*()"}
	var number = []string{"1", "12", "123", "1234", "12345", "123456", "1234567", "12345678",
		"123456789", "987654321", "87654321", "7654321", "654321", "54321", "4321",
		"321", "21", "6s", "66", "666", "6666", "66666", "666666", "6666666", "66666666",
		"8", "88", "888", "8888", "88888", "888888", "8888888", "88888888", "0", "00",
		"000", "0000", "00000", "000000", "0000000", "00000000"}

	for i := 0; i < len(char); i++ {
		fmt.Println(char[i])

	}
	for j := 0; j < len(number); j++ {
		fmt.Println(number[j])
	}
	domain := os.Args[1]
	fmt.Println(domain)
}