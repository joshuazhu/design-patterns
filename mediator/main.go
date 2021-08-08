package main

func main() {
	sm := newStationManger()
	pt := passengerTrain{mediator: sm}
	ft := freightTrain{mediator: sm}

	pt.arrive()
	ft.arrive()
	pt.depart()
	ft.depart()
}
