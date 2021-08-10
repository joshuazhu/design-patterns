package main

type circle struct {
	side int
}

func (c circle) accept(v visitor) {
	v.visitCircle(c)
}
