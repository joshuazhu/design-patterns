package main

import "fmt"

type Director struct {
	carBuilder Builder
}

func (d *Director) setCarBuilder(cb Builder) {
	d.carBuilder = cb
}

func (d *Director) makeSUV() {
	fmt.Println("Making SUV")
	d.carBuilder.setEngine("SUV engine")
	d.carBuilder.setSeats(4)
	d.carBuilder.setGPS(true)
	d.carBuilder.setTripComputer(true)
}

func (d *Director) makeSportsCar() {
	fmt.Println("Making Sports Car")
	d.carBuilder.setEngine("Sports engine")
	d.carBuilder.setSeats(2)
	d.carBuilder.setGPS(true)
	d.carBuilder.setTripComputer(true)
}
