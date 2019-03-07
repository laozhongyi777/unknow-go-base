package main

import "fmt"

func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1]
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD"
	fmt.Println(a, ok) //每一层都要make初始化，不然会报错
}
