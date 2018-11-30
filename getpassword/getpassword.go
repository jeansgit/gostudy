package main

import "fmt"
import "os"
import "strings"
import "strconv"

func main() {

	var char = []string{"", "`", "~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "=", "_",
		"+", ",", ".", "/", ";", "\\", "\"", "<", ">", "?", "|", ":", "[", "]", "{", "}",
		"!@", "!#", "@#", "@!", "#@", "#!", "!@#", "!#@", "#@!", "#!@", "@!#", "@#!", "!@#$",
		"!@#$%", "!@#$%^", "!@#$%^&", "!@#$%^&*", "!@#$%^&*(", "!@#$%^&*()"}
	var number = []string{"1", "12", "123", "1234", "12345", "123456", "1234567", "12345678",
		"123456789", "987654321", "87654321", "7654321", "654321", "54321", "4321",
		"321", "21", "6", "66", "666", "6666", "66666", "666666", "6666666", "66666666",
		"8", "88", "888", "8888", "88888", "888888", "8888888", "88888888", "0", "00",
		"000", "0000", "00000", "000000", "0000000", "00000000", "520", "woaini", "1314"}
	/*
		for i := 0; i < len(char); i++ {
			fmt.Println(char[i])

		}
		for j := 0; j < len(number); j++ {
			fmt.Println(number[j])
		}
	*/
	domain := os.Args[1]
	thepath := os.Args[2]
	file := filecreate(thepath)
	defer closefile(file)
	fmt.Println(domain)
	str := strings.Split(domain, ".")
	fmt.Println(str[0])
	for y := 1970; y < 2019; y++ {
		fmt.Println(str[0] + strconv.Itoa(y))       // 生成密码baidu2018
		filewrite(file, (str[0] + strconv.Itoa(y))) //写入文件
		fmt.Println(strconv.Itoa(y) + str[0])       // 生成密码2018baidu
		filewrite(file, (strconv.Itoa(y) + str[0])) //写入文件
		for _, m := range char {
			fmt.Println(str[0] + m + strconv.Itoa(y))       // 生成密码baidu@2018
			filewrite(file, (str[0] + m + strconv.Itoa(y))) //写入文件
			fmt.Println(strconv.Itoa(y) + m + str[0])       // 生成密码2018@baidu
			filewrite(file, (strconv.Itoa(y) + m + str[0])) //写入文件
		}
	}
	for _, m := range number {
		fmt.Println(str[0] + m)       // 生成密码baidu123
		filewrite(file, (str[0] + m)) //写入文件
		fmt.Println(m + str[0])       // 生成密码123baidu
		filewrite(file, (m + str[0])) //写入文件
	}
	for _, thechar := range char {
		for _, m := range number {
			fmt.Println(str[0] + thechar + m)       // 生成密码baidu@123
			filewrite(file, (str[0] + thechar + m)) //写入文件
			fmt.Println(m + thechar + str[0])       // 生成密码123@baidu
			filewrite(file, (m + thechar + str[0])) //写入文件
		}
	}
}
func filecreate(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)

	}
	return file
}
func filewrite(filename *os.File, str string) {
	fmt.Fprintln(filename, str)
}
func closefile(filename *os.File) {
	filename.Close()
}
