package domain

import "errors"

var (
	ErrParkingFull  = errors.New("parking lot is full")
	ErrCarNotFound  = errors.New("car not found")
)

type ParkingLotManager struct {
	Lots []*ParkingLot
}

func NewParkingLotManager(capacity int) *ParkingLotManager {
	lots := make([]*ParkingLot, capacity)

	for i := 0; i < capacity; i++ {
		lots[i] = NewParkingLot(i + 1)
	}
	return &ParkingLotManager{Lots: lots}
}

func (m *ParkingLotManager) Park(car *Car) (*ParkingLot, error) {
	for _, lot := range m.Lots {
		if lot.IsEmpty() {
			lot.Park(car)
			return lot, nil
		}
	}
	return nil, ErrParkingFull
}

func (m *ParkingLotManager) Leave(carNumber string) (*ParkingLot, *Car, error) {
	for _, lot := range m.Lots {
		if !lot.IsEmpty() && lot.Car.CarNumber == carNumber {
			car := lot.Leave()
			return lot, car, nil
		}
	}
	return nil, nil, ErrCarNotFound
}
