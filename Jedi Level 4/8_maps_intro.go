package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	// get James' age
	fmt.Println(m["James"])

	// Check if a variable actually exists as it returns zeros... comma ok idiom
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	// A more concise way of doing this is
	if v, ok := m["James"]; ok {
		fmt.Println("This is the if print ", v)
	}
}
