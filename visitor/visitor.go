package main

type visitor interface {
	visitSquare(square)
	visitRectangle(rectangle)
	visitCircle(circle)
}
