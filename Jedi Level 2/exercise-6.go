package main

import "fmt"

const (
	a int = 2019 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
