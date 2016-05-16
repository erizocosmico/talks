import "fmt"

// All Animals should have a method named Speak that returns a string
type Animal interface {
	Speak() string
}

type Human struct {
	Name string
}

// Now Human IS an Animal
func (h Human) Speak() string {
	return fmt.Sprintf("my name is %s", h.Name)
}
