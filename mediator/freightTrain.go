package main

import "fmt"

type freightTrain struct {
	mediator mediator
}

func (p *freightTrain) arrive() {
	if p.mediator.canArrive(p) {
		fmt.Println("Freight train: Arrived")
		return
	}

	fmt.Println("Freight train: Blocked")
}

func (p *freightTrain) depart() {
	fmt.Println("Freight train: Leaving station")
	p.mediator.notifyDeparture()
}

func (p *freightTrain) permitArrival() {
	fmt.Println("Freight train: Arrival permitted")
	p.arrive()
}
