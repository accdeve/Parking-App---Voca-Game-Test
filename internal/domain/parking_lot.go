package domain

type ParkingLot struct {
	SlotNumber int
	Car        *Car
}

func NewParkingLot(SlotNumber int) *ParkingLot {
	return &ParkingLot{
		SlotNumber: SlotNumber,
		Car:        nil,
	}
}

func (p *ParkingLot) Park(car *Car) {
	p.Car = car
}

func (p *ParkingLot) Leave() *Car {
	car := p.Car
	p.Car = nil
	return car
}

func (p *ParkingLot) IsEmpty() bool {
	return p.Car == nil
}
