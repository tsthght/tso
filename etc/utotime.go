package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(1596099179, 0)
	fmt.Println(t.String())
	fmt.Printf("\n")
	t = time.Unix(1596099845, 0)
	fmt.Println(t.String())
	fmt.Printf("\n")
	t = time.Unix(1596099179, 0)
	fmt.Println(t.String())
	fmt.Printf("\n")
	t = time.Unix(1596099845, 0)
	fmt.Printf(t.String())
	fmt.Printf("\n")
	t = time.Unix(1596099845, 0)
	fmt.Printf(t.String())
}
