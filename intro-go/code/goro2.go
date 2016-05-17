package main

import (
	"fmt"
	"time"
)

func plusOne(n int) {
	time.Sleep(1 * time.Second)
	fmt.Println(n + 1)
}

func main() {
	for i := 0; i < 5; i++ {
		go plusOne(i)
	}
}
