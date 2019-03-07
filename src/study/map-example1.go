package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "OK"
	delete(m, 1) //删除键值
	a := m[1]
	fmt.Println(a)
}
