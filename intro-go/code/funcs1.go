package main

import "fmt"

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(swap(42, 13))
}
