package main

func main() {
	r := rectangle{with: 2, long: 4}
	c := circle{side: 2}
	s := square{side: 3}

	ac := areaCalculator{}
	pc := perimeterCalculator{}

	r.accept(ac)
	r.accept(pc)

	c.accept(ac)
	c.accept(pc)

	s.accept(ac)
	s.accept(pc)
}
