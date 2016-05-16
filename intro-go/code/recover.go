package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic!")
		}
	}()
	panic("foo")
	fmt.Println("END")
}
