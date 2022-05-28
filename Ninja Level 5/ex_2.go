package main

import (
	"fmt"
)

type person struct {
	first_name                  string
	last_name                   string
	favorite_ice_cream_flavours []string
}

func main() {
	p1 := person{
		first_name:                  "Job",
		last_name:                   "Ballard",
		favorite_ice_cream_flavours: []string{"Rum n' Raisin", "Mint", "Vanilla"},
	}

	p2 := person{
		first_name:                  "Keke",
		last_name:                   "Blue",
		favorite_ice_cream_flavours: []string{"Chocolate", "Cookies n' Cream", "Pistachio"},
	}

	m := map[string]person{}
	m[p1.first_name] = p1
	m[p2.first_name] = p2

	for _, v := range m {
		fmt.Println(v.first_name, v.last_name)
		for i, val := range v.favorite_ice_cream_flavours {
			fmt.Println(i, val)
		}
		fmt.Println("------------")
	}
}
