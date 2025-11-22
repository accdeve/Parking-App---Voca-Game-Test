package domain

import "time"

type Car struct{
	CarNumber string
	ParkedAt time.Time
}

func NewCar(carNumber string) *Car{
	return &Car{
		CarNumber: carNumber,
		ParkedAt: time.Now(),
	}
}