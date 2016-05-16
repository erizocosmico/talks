package main

import "fmt"

func main() {
	var i = 5
	if i == 6 {
		fmt.Println("i was 6")
	} else if i == 7 {
		fmt.Println("i was 7")
	} else {
		fmt.Printf("i was %v", i)
	}
}
