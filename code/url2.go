package main

import "fmt"
import "net/url"
import "strings"

func main() {
	s := "https://mbd.baidu.com/newspage/data/landingsuper?context=%7B\"nid\"%3A\"news_10014101464176517173\"%7D&n_type=0&p_from=1"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	//fmt.Println(h[1])
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	//fmt.Println(m["k"][0])
}
