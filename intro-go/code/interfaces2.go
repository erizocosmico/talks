package main

import "fmt"

type Animal interface {
	Speak() string
}

type Human struct {
	Name string
}

func (h Human) Speak() string {
	return fmt.Sprintf("my name is %s", h.Name)
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meoooooow"
}

func main() {
	var animals = []Animal{
		Human{"Steve"},
		Cat{},
		Human{"Adam"},
	}

	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}
