package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9, 10, 1, 1, 1, 1}
	copy(s2[2:4], s1[1:3]) //m没怎么懂
	fmt.Println(s2)
}
