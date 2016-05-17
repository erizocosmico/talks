package main

import "fmt"

func main() {
	var foo interface{} = 42
	var bar int = foo.(int)
	fmt.Println(bar)

	var baz interface{} = "hello"
	num, ok := baz.(int)
	if !ok {
		fmt.Println("baz is not a number!")
	} else {
		fmt.Println("baz is", num)
	}
}
