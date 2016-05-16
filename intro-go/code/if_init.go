package main

import "fmt"

func main() {
	if err := canFail(); err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("not failed")
	}
}
