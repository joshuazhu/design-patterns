package main

import "fmt"

type CarBuilder struct {
	car *Car
}

func newCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: &Car{},
	}
}

func (cb *CarBuilder) reset() {
	cb.car = &Car{}
}

func (cb *CarBuilder) setSeats(seats int) {
	fmt.Println("Setting seats via car builder")
	cb.car.numberOfSeats = seats
}

func (cb *CarBuilder) setEngine(engineType string) {
	fmt.Println("Setting engine type via car builder")
	cb.car.engineType = engineType
}

func (cb *CarBuilder) setTripComputer(hasTripComputer bool) {
	fmt.Println("Setting trip computer via car builder")
	cb.car.tripComputer = hasTripComputer
}

func (cb *CarBuilder) setGPS(hasGPS bool) {
	fmt.Println("Setting gps via car builder")
	cb.car.gps = hasGPS
}

func (cb *CarBuilder) getCar() *Car {
	c := cb.car
	cb.reset()
	return c
}
