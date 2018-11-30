package main

import "fmt"

//import "crypto/sha1"
import "crypto/md5"

func main() {
	str := "123456"
	md5hash := md5.New()
	text := []byte(str)
	md5hash.Write(text)
	fmt.Printf("%x %x\n", md5hash.Sum(nil), md5hash.Sum(text))
	md5hash.Write(nil)
	fmt.Printf("%x %x\n", md5hash.Sum(nil), md5hash.Sum(text))
}
