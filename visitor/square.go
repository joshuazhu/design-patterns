package main

type square struct {
	side int
}

func (s square) accept(v visitor) {
	v.visitSquare(s)
}
