package main

type Builder interface {
	reset()
	setSeats(int)
	setEngine(string)
	setTripComputer(bool)
	setGPS(bool)
}
