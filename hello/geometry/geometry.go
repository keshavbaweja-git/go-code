package geometry

// Geometry models geometric shapes
type Geometry interface {
	Area() float64
	Perimeter() float64
}

// Rect models a rectangle
type Rect struct {
	L, B float64
}

// Area returns area of a rectangle
func (r Rect) Area() float64 {
	return r.L * r.B
}

// Perimeter returns perimeter of a rectangle
func (r Rect) Perimeter() float64 {
	return 2*r.L + 2*r.B
}
