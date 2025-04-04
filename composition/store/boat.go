package store

// простой пример использования композии
type Boat struct {
	*Product
	Capacity  int
	Motorized bool
}

// простой пример использования композии
func NewBoat(name string, price float64, capacity int,	motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
	}
}
