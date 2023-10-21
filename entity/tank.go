package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	TankCapacityError = errors.New("invalid tank capacity, value must be between 0 and 2000")
)

type Tank struct {
	Base
	capacity float64
	quantity float64
}

func NewTank(capacity float64) (Tank, error) {
	err := validateNewTank(capacity)
	if err != nil {
		return Tank{}, err
	}
	return Tank{
		Base: Base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		capacity: capacity,
	}, nil
}

func validateNewTank(capacity float64) error {
	if capacity <= 0 || capacity > 2000 {
		return TankCapacityError
	}
	return nil
}
