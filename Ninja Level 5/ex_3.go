package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	f150 := truck{
		vehicle:   vehicle{doors: 4, color: "Blue"},
		fourWheel: true,
	}
	impala := sedan{
		vehicle: vehicle{doors: 4, color: "White"},
		luxury:  false,
	}

	fmt.Println(f150)
	fmt.Println(impala)

	fmt.Println(f150.color)
	fmt.Println(impala.color)

}
