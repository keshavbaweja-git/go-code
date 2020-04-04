package main

import (
	"fmt"

	"github.com/keshavbaweja-git/hello/morestrings"
	"github.com/keshavbaweja-git/hello/rect"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	rect1 := rect.Rect{L: 10, B: 5}
	fmt.Println(rect1.Area())
}
