package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		println(i, v)
	}
	fmt.Printf("%T\n", x)

	s1 := x[0:4]

	fmt.Println(s1)
}
