package main

import "fmt"

// Defer pushes the defered function to the end of the main() series
// of actions. It will run after the last function runs

func main() {
	defer foo()
	bar()
}

func foo() { fmt.Println("foo") }
func bar() { fmt.Println("bar") }
