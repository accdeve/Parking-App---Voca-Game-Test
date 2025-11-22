package memory

import "parkee/internal/domain"

type MemoryParkingRepository struct{
	manager *domain.ParkingLotManager
}

func NewMemoryParkingRepository() *MemoryParkingRepository{
	return &MemoryParkingRepository{}
}

func (r *MemoryParkingRepository) Save(m *domain.ParkingLotManager) error{
	r.manager = m
	return nil
}

func (r *MemoryParkingRepository) Load() (*domain.ParkingLotManager, error){
	return r.manager, nil
}