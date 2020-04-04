package rect

// Rect models a rectangle
type Rect struct {
	L, B int
}

// Area returns area of a rectangle
func (r Rect) Area() int {
	return r.L * r.B
}

// Perimeter returns perimeter of a rectangle
func (r Rect) Perimeter() int {
	return 2*r.L + 2*r.B
}
