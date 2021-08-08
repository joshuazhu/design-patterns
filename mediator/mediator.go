package main

type mediator interface {
	canArrive(t train) bool
	notifyDeparture()
}
