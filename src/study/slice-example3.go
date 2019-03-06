package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 3, 6)
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
	s1 = append(s1, 2, 3, 4)
	fmt.Println(s1) //重新分配内存地址
}
