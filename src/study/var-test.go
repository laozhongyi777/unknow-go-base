package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 65 //string　　类型是　６５个字节即大写的A
	//b := string(a)
	b := strconv.Itoa(a)
	fmt.Println(b)

}
