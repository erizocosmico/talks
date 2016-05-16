package main

import "fmt"

func main() {
	a := func() string {
		return "foo"
	}
	fmt.Println(a())
}
