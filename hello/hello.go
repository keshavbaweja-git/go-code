package main

import (
	"fmt"

	"github.com/keshavbaweja-git/hello/carpurchase"
	"github.com/keshavbaweja-git/hello/geometry"
)

func main() {
	fmt.Println(carpurchase.NbMonths(2000, 8000, 1000, 1.5))
}

func printGeometry(g geometry.Geometry) {
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}
