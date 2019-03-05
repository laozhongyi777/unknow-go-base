package main

import (
	"fmt"
)

func main() {
LABLE:
	for i := 0; i < 10; i++ {
		if i > 2 {
			break LABLE
		} else {
			fmt.Println(i)
		}
	}
}
