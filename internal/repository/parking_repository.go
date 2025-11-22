package repository

import "parkee/internal/domain"

type ParkingRepository interface{
	Save(manager *domain.ParkingLotManager) error
	Load()(*domain.ParkingLotManager, error)
}