package main

import "fmt"

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	PB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(PB)

}
