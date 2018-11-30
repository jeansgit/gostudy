package main
import "fmt"
func main() {
	var countrymap map[string]string  //声明变量 

	//初始化集合 创建一个nil集合 ，nil map 不能用来存放键值对
	countrymap = make(map[string]string)

	countrymap["1"]="aaa"
	countrymap["2"]="bbb"
	countrymap["3"]="ccc"
	countrymap["4"]="ddd"

	for i:=range countrymap{
		fmt.Println(countrymap[i])
	}

}