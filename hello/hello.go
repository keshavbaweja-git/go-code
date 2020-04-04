package main

import (
	"fmt"

	"github.com/keshavbaweja-git/hello/geometry"
	"github.com/keshavbaweja-git/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	rect1 := geometry.Rect{L: 10, B: 5}
	printGeometry(rect1)
}

func printGeometry(g geometry.Geometry) {
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}
