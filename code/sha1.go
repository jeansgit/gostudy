package main

import "fmt"
import "crypto/sha1"
import "crypto/md5"

func main() {
	text := "123456"
	sha1res := sha1.New()       //产生sha1的方式
	md5res := md5.New()         //产生md5的方式
	sha1res.Write([]byte(text)) //写入要加密的字符串并进行处理,这里是字符串 所以用byte()

	/*这个用来得到最终的散列值的字符切片。
	Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。*/
	sha1result := sha1res.Sum(nil)

	md5res.Write([]byte(text))
	md5result := md5res.Sum(nil)
	fmt.Println(sha1res)
	fmt.Printf("%x\n", sha1result) //以16进制格式输出加密结果
	fmt.Println(md5res)
	fmt.Printf("%x\n", md5result) //以16进制格式输出加密结果
}
