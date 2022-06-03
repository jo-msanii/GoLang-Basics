package main

import "fmt"

func main() {
	foo()
	bar("james")
	s1 := woo("Monerpenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// Takes no  arguments returns nothing
func foo() {
	fmt.Println("hello from woo!")
}

// Takes an argument returns nothing
func bar(s1 string) {
	fmt.Println("Hello,", s1)
}

// Takes and argument and returns something
func woo(s2 string) string {
	return fmt.Sprint("Hello from woo, ", s2)
}

// Takes and argument and has multiple returns
func mouse(s3 string, s4 string) (string, bool) {
	a := fmt.Sprint(s3, s4, `, says "Hello"`)
	b := false
	return a, b
}
