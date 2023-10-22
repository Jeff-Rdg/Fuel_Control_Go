package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	TankCapacityError       = errors.New("invalid tank capacity, value must be between 0 and 2000")
	TankRefuelDateError     = errors.New("invalid date to refuel")
	TankRefuelQuantityError = errors.New("quantity in the tank plus the amount of supply exceeds tank capacity")
)

type Tank struct {
	Base
	capacity float64
	quantity float64
	history  []FuelTank
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

func (tank *Tank) Refuel(fuel Fuel, date time.Time) error {
	if fuel.quantity > tank.capacity {
		return TankCapacityError
	}
	if tank.quantity+fuel.quantity > tank.capacity {
		return TankRefuelQuantityError
	}
	if date.After(time.Now()) {
		return TankRefuelDateError
	}
	fuelTank, err := NewFuelTank(fuel.id, tank.id, fuel.quantity, date)
	if err != nil {
		return err
	}

	tank.quantity += fuel.quantity
	tank.history = append(tank.history, fuelTank)

	return nil
}

func (tank *Tank) GetId() uuid.UUID {
	return tank.id
}

func (tank *Tank) GetCapacity() float64 {
	return tank.capacity
}

func (tank *Tank) SetCapacity(value float64) {
	if tank.capacity != value && value > 0 && value <= 2000 {
		tank.capacity = value
	}
}

func (tank *Tank) GetQuantity() float64 {
	return tank.quantity
}

func (tank *Tank) SetQuantity(value float64) {
	if tank.quantity != value && value > 0 && value <= 2000 {
		tank.quantity = value
	}
}
