package main

import (
	"fmt"
)

func main() {
	//	var a [2]int  //为赋值默认为　０
	//  a :=[...]int{1, 2, 3, 4, 5} 使用类型推断，　三个点表示n个元素，等于右边包含值的个数
	var a [2]int
	var b [2]int
	b = a // 只能为等于或者不等于

	fmt.Println(b)
}
