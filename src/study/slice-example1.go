package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	//	s1 := a[5:]  //取第6个元素到最后一个
	s1 := a[5:10] //元素是从位数是从０　开始的打印　a[ 6 7 8 9 10]
	fmt.Println(s1)
}
