package main

import "fmt"

type Vehicle struct {
	Seats int
	Speed float64
}

type Car struct {
	Wheels int
	Color  string
	Vehicle
}

func main() {
	car := Car{
		Wheels: 4,
		Color:  "red",
		Vehicle: Vehicle{
			Seats: 5,
			Speed: 65.,
		},
	}
	fmt.Println(car)

	car.Speed = 67.
	fmt.Println(car)
}
