package main

import "fmt"

type Car struct {
	Wheels int
	Speed  float64
	Color  string
}

func main() {
	// is empty
	var car *Car
	fmt.Println(car)

	car = &Car{
		Wheels: 4,
		Speed:  45.,
		Color:  "green",
	}
	fmt.Println(car)

	car.Speed = 50.
	fmt.Println(car)
}
