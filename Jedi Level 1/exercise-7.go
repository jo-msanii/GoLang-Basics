package main

import "fmt"

func main() {

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\t\t\t\t%x\n", kb, kb, kb)
}
