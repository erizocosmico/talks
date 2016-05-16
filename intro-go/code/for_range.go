package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	for i, num := range numbers {
		fmt.Println(i, num)
	}
}
