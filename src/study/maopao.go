package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	num := len(a)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[j]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}
