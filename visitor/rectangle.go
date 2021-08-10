package main

type rectangle struct {
	with int
	long int
}

func (r rectangle) accept(v visitor) {
	v.visitRectangle(r)
}
