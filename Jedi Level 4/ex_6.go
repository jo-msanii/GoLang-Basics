package main

import "fmt"

func main() {
	src_states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	// fmt.Println("src_states has a length of: ", len(src_states))

	states_slice := make([]string, 50, 50)

	for i := 0; i <= 49; i++ {
		states_slice[i] = src_states[i]
	}

	fmt.Println("Length of my slice is ", len(states_slice))
	fmt.Println("Capacity of my slice is ", cap(states_slice))

	for i := 0; i <= 49; i++ {
		fmt.Println(i, "\t", states_slice[i])
	}

	// for {
	// 	if counter <
	// 			states_slice[counter] = src_states[counter]
	// 			fmt.Println(states_slice[counter], src_states[counter])
	// 			counter++
	// 		}

	// fmt.Println(states_slice)
	// fmt.Println(src_states[0])
	// fmt.Printf("%T", states_slice)
	// fmt.Printf("%T", src_states)

	// states_slice[0] = src_states[0]
	// fmt.Println(states_slice[0])

}
