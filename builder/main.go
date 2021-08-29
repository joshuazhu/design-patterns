package main

func main() {
	d := Director{}
	cb := newCarBuilder()

	d.setCarBuilder(cb)
	d.makeSUV()

	cmb := newCarManualBuilder()
	d.setCarBuilder(cmb)
	d.makeSportsCar()
}
