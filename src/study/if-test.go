package main

import "fmt"

func main() {
	a := 10
	if a, b := 1, 2; a > b {
		fmt.Println(a) //内部条件被隐藏
	}
	fmt.Println(a)
}
