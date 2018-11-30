package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)

	daytime := time.Date(2001, 11, 2, 21, 23, 22, 6665326, time.UTC)
	p(daytime)
	p(daytime.Year())
	p(daytime.Weekday())

}
