package main

import "fmt"

func main() {
	A(1, 2, 3, 4, 5, 6, 7)
}

func A(a ...int) int { //...代表不定长变参
	fmt.Println(a)
}
