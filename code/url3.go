package main

import "fmt"
import "net/url"
import "strings"

func main() {
	testurl := "http://222.76.242.215:8586/cas-server/login?service=http%3A%2F%2F222.76.242.215%3A8585%2Fzhhy"
	u, err := url.Parse(testurl)
	if err != nil {
		panic(err)
	}
	f := fmt.Println
	f(u)
	f(u.Scheme)
	f(u.Host)
	str := strings.Split(u.Host, ":")
	f(str[0])
	f(str[1])
	f(u.Path)
	f(u.Fragment)
	f(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	f(m)
	f(m["service"][0])
}
