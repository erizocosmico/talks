type Speaker interface {
	Speak() string
}

type Animal interface {
	Speaker
}
