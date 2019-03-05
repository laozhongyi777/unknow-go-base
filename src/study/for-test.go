package main

import "fmt"

//ex.. 1
//func main()  {
//	a :=1
//	for {
//		a++
//		if a >3 {
//			break
//		}
//	}

//ex... 2
//	fmt.Println(a)
//}
//func main()  {
//	a :=1
//	for a <= 3  {
//		a ++
//	}
//	fmt.Println(a)
//}

func main() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
	}
	fmt.Println(a)
}
