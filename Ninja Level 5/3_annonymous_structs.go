package main

import "fmt"

// Annonymous is unnamed. So you can declare and use it
// if it is a one off use-case scenario.

// Useful to avoid code pollution where you don't have repeated
// use of a struct

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
