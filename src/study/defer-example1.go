package main

import "fmt"

func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c") //倒序执行　　先定义，后执行
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) //闭包编程地址直接引用３
		}()
	}

}
