package service

import (
	"parkee/internal/domain"
	"parkee/internal/repository"
)

type ParkingService struct {
	repo repository.ParkingRepository
}

func NewParkingService(repo repository.ParkingRepository) *ParkingService {
	return &ParkingService{repo: repo}
}

func (s *ParkingService) CreateParkingLot(capacity int) error {
	manager := domain.NewParkingLotManager(capacity)
	return s.repo.Save(manager)
}

func (s *ParkingService) Park(carNumber string) (*domain.ParkingLot, error) {
	manager, _ := s.repo.Load()

	car := domain.NewCar(carNumber)
	lot, err := manager.Park(car)
	if err != nil {
		return nil, err
	}

	s.repo.Save(manager)
	return lot, nil
}

func (s *ParkingService) Leave(carNumber string, hours int) (*domain.ParkingLot, *domain.Car, int, error) {
	manager, _ := s.repo.Load()

	lot, car, err := manager.Leave(carNumber)
	if err != nil {
		return nil, nil, 0, err
	}

	// Hitung charge berdasarkan aturan PDF
	charge := calculateCharge(hours)

	s.repo.Save(manager)
	return lot, car, charge, nil
}

func (s *ParkingService) Status() []*domain.ParkingLot {
	manager, _ := s.repo.Load()
	return manager.Lots
}

func calculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + (hours-2)*10
}
