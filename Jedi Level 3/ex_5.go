package main

import (
	"fmt"
)

func main() {

	// Print out the remainder (modulus) which is
	// found for each number betwen 10 and 100 when
	// it is divided by 4

	for number := 10; number < 101; number++ {
		if number%4 != 0 {
			fmt.Println("For ", number, " the remainder is ", number%4)
		}
	}

	fmt.Println("end")

}
