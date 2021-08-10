package main

import (
	"fmt"
	"math"
)

type areaCalculator struct{}

func (a areaCalculator) visitRectangle(r rectangle) {
	fmt.Printf("Rectangle area: %v\n", r.long*r.with)
}

func (a areaCalculator) visitSquare(s square) {
	fmt.Printf("Square area: %v\n", s.side*s.side)
}

func (a areaCalculator) visitCircle(c circle) {
	fmt.Printf("Circle area: %v\n", math.Pi*float64(c.side*c.side))
}
