package main

import "fmt"

type CarManualBuilder struct {
	car *Car
}

func newCarManualBuilder() *CarManualBuilder {
	return &CarManualBuilder{
		car: &Car{},
	}
}

func (cb *CarManualBuilder) reset() {
	cb.car = &Car{}
}

func (cb *CarManualBuilder) setSeats(seats int) {
	fmt.Println("Setting seats via car manual builder")
	cb.car.numberOfSeats = seats
}

func (cb *CarManualBuilder) setEngine(engineType string) {
	fmt.Println("Setting engine type via car manual builder")
	cb.car.engineType = engineType
}

func (cb *CarManualBuilder) setTripComputer(hasTripComputer bool) {
	fmt.Println("Setting trip computer via car manual builder")
	cb.car.tripComputer = hasTripComputer
}

func (cb *CarManualBuilder) setGPS(hasGPS bool) {
	fmt.Println("Setting gps via car manual builder")
	cb.car.gps = hasGPS
}

func (cb *CarManualBuilder) getCar() *Car {
	c := cb.car
	cb.reset()
	return c
}
