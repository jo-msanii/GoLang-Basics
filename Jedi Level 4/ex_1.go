package main

import "fmt"

func main() {
	// Compisite literal means we determine the size up-front?
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("A is %T", a)
}
