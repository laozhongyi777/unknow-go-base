package main

import (
	"fmt"
)

func main() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			//continue LABEL  替换成goto看效果
			goto LABEL
		}
	}
}
