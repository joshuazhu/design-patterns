package main

type stationManager struct {
	isStationFree bool
	trains        []train
}

func newStationManger() *stationManager {
	return &stationManager{
		isStationFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isStationFree {
		s.isStationFree = false
		return true
	}

	s.trains = append(s.trains, t)
	return false
}

func (s *stationManager) notifyDeparture() {
	if len(s.trains) > 0 {
		t := s.trains[0]

		s.trains = s.trains[1:]
		s.isStationFree = true

		t.permitArrival()
	}
}
