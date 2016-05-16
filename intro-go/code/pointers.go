package main

import "fmt"

func main() {
	var p *int
	myInt := 3
	p = &myInt

	fmt.Printf("p = %v, *p = %v", p, *p)

	*p = 5
	fmt.Printf("\n\np = %v, *p = %v, myInt = %v", p, *p, myInt)
}
