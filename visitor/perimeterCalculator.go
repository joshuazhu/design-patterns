package main

import (
	"fmt"
	"math"
)

type perimeterCalculator struct{}

func (a perimeterCalculator) visitRectangle(r rectangle) {
	fmt.Printf("Rectangle perimeter: %v\n", (r.long+r.with)*2)
}

func (a perimeterCalculator) visitSquare(s square) {
	fmt.Printf("Square perimeter: %v\n", s.side*4)
}

func (a perimeterCalculator) visitCircle(c circle) {
	fmt.Printf("Circle perimeter: %v\n", math.Pi*float64(c.side*2))
}
