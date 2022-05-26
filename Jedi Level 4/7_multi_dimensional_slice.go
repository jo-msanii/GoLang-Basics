package main

import "fmt"

func main() {
	/*
		Creating a slice of a slice
	*/

	jb := []string{"Rose", "Susan", "Tilda", "Edward", "David"}
	ro := []string{"Veronica", "Keke", "Edwin", "Uche", "Krystal"}
	fmt.Println(jb)
	fmt.Println(ro)

	xp := [][]string{jb, ro}
	fmt.Println(xp)
}
