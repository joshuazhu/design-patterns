package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) arrive() {
	if p.mediator.canArrive(p) {
		fmt.Println("Passenger train: Arrived")
		return
	}

	fmt.Println("Passenger train: Blocked")
}

func (p *passengerTrain) depart() {
	fmt.Println("Passenger train: Leaving station")
	p.mediator.notifyDeparture()
}

func (p *passengerTrain) permitArrival() {
	fmt.Print("Passenger train: Arrival permitted")
	p.arrive()
}
