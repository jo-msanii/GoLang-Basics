package main

import "fmt"

func main() {

	type MyVar int
	const a int = 1
	const b MyVar = 2

	fmt.Println(a, b)
}
