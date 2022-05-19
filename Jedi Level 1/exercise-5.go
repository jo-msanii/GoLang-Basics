package main

import "fmt"

const (
	a = 42
	b = 42.78
	c = "james Bond"

	e = iota
	f = iota
	g = iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}
