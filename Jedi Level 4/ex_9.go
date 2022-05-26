package main

import "fmt"

func main() {
	// Create a map key last_first_name with the values being a []string
	// slice with three things liked
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["ballard_job"] = []string{"Coding", "Motorcycles", "Keke"}

	// Get all keys into []string slice
	// Range over the map
	record_numbers := 1
	for k, v := range m {
		fmt.Printf("RECORD %v: %v\n", record_numbers, k)
		for k2, v2 := range v {
			fmt.Printf("\t %v %v\n", k2, v2)
		}
		record_numbers++
	}

}
