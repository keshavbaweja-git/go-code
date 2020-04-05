package main

import (
	"fmt"

	"github.com/keshavbaweja-git/hello/geometry"
	"github.com/keshavbaweja-git/hello/simplestringindicies"
)

func main() {
	fmt.Println(simplestringindicies.Solution("(abc)", 0))
}

func printGeometry(g geometry.Geometry) {
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}
