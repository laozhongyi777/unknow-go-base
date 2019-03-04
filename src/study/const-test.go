package main

import "fmt"

const (
	B float64 = 1 << (iota * 10)
	KB
	BM
	GB
	PB
)

func main() {
	fmt.Println(PB)
}
