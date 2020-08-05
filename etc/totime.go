package main

import (
	"fmt"
	"time"
)

func main() {
	ts1 := int64(418515433628106755)
	t1 := TSOToRoughTime(ts1)

	time.Now().Unix()

	fmt.Println(t1.String())
}

// TSOToRoughTime translates tso to rough time that used to display
func TSOToRoughTime(ts int64) time.Time {
	t := time.Unix(ts>>18/1000, 0)
	return t
}
