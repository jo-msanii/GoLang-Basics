package main

import "fmt"

func main() {

	// Print all the numbers between 0 and 10,000

	w := 65
	for x := 65; x < 90; x++ {
		fmt.Println(w)
		for z := 0; z < 3; z++ {
			fmt.Printf("\t%#U \n", x)
		}
		w++
	}
	fmt.Println("end")

}
