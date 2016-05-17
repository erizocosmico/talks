package main

import "fmt"

func main() {
	var numbers = make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			numbers <- i + 1
		}
		close(numbers) // <- we are closing the channel
	}()

	n := <-numbers
	fmt.Println("FIRST", n)

	for n := range numbers {
		fmt.Println(n)
	}
}
