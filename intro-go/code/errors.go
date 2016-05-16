package main

import (
	"errors"
	"fmt"
)

func canFail() (int, error) {
	return 0, errors.New("oops, I failed")
}

func main() {
	v, err := canFail()
	if err != nil {
		fmt.Println("failed :(")
	} else {
		fmt.Println(v)
	}
}
