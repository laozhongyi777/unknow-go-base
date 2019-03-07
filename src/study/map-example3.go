package main

import "fmt"

func main() {
	sm := make([]map[int]string, 5) //定义元素个数为５
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
