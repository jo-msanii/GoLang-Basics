package main

import "fmt"

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

	fmt.Println("Customer Details:")
	fmt.Println("\t", p1.first_name, p1.last_name)
	fmt.Println("Favourite Flavors:")
	for k, v := range p1.favorite_ice_cream_flavours {
		fmt.Println("\t", k, v)
	}

	fmt.Println("\n")

	fmt.Println("Customer Details:")
	fmt.Println("\t", p2.first_name, p2.last_name)
	fmt.Println("Favourite Flavors:")
	for k, v := range p2.favorite_ice_cream_flavours {
		fmt.Println("\t", k, v)
	}
}
