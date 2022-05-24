package main

import (
	"fmt"
	"time"
)

func main() {

	// Print out the years you have been alive
	year := time.Now()
	birth_year := 1984
	year_for_loop := year.Year()
	for {
		fmt.Println(year_for_loop)
		year_for_loop--
		if year_for_loop <= birth_year {
			break
		}
	}

	fmt.Println("end")

}
