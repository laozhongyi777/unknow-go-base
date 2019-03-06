package main

import (
	"fmt"
)

//func main()  {
//	s1 :=make([]int,3 ,40)
//	fmt.Println(len(s1),cap(s1))
//}
func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j'}
	sa := a[2:5]
	sb := sa[1:3] //slice的元素中取，元素不能超过slice的元素容量
	fmt.Println(string(sb))
}
