package main

import (
	"fmt"
	"time"
)

func main() {
	str := "2020-07-30 16:30:00"
	loc, _ := time.LoadLocation("Local")
	cur, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	//cur := time.Now()
	fmt.Printf("cur: %v\n", cur)
	t := cur.Unix() * 1000
	tso := t << 18
	fmt.Printf("tso: %v\n", tso);
}
