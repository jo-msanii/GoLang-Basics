package main

import (
	"fmt"
)

func main() {

	job := struct {
		alive bool
		sex   int
		race  string
	}{
		alive: true,
		sex:   1,
		race:  "Black",
	}

	fmt.Println(job)
}
