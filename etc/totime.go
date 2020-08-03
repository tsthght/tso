package main

import (
	"fmt"
	"time"
)

func main() {
	ts1 := int64(1596098534043317)
	t1 := TSOToRoughTime(ts1)

	fmt.Println(t1.String())
}

// TSOToRoughTime translates tso to rough time that used to display
func TSOToRoughTime(ts int64) time.Time {
	t := time.Unix(ts>>18/1000, 0)
	return t
}
