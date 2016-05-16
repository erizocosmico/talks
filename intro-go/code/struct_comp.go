type Vehicle struct {
	Seats int
	Speed float64
}

type Car struct {
	Wheels int
	Color  string
	Vehicle
}
