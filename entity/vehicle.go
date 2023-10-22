package entity

import (
	"FuelControl/entity/enum"
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	VehiclePlateError    = errors.New("invalid plate")
	VehicleOdometerError = errors.New("invalid odometer")
	VehicleTypeError     = errors.New("invalid VehicleType, valid types: " + enum.GetVehicleTypes())
	VehicleOwnerError    = errors.New("invalid Owner")
)

type Vehicle struct {
	Base
	plate       string
	odometer    int
	vehicleType enum.VehicleType
	Owner
	history []VehicleTank
}

func NewVehicle(plate string, odometer int, vehicleType string, owner Owner) (Vehicle, error) {
	err := validateVehicle(plate, odometer, vehicleType, owner)
	if err != nil {
		return Vehicle{}, err
	}
	return Vehicle{
		Base: Base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		plate:       plate,
		odometer:    odometer,
		vehicleType: enum.VehicleType(vehicleType),
		Owner:       owner,
	}, nil
}

func validateVehicle(plate string, odometer int, vehicleType string, owner Owner) error {
	if plate == "" {
		return VehiclePlateError
	}
	if odometer < 0 {
		return VehicleOdometerError
	}
	if !enum.ValidVehicleType(vehicleType) {
		return VehicleTypeError
	}

	if (Owner{} == owner) {
		return VehicleOwnerError
	}
	return nil
}

func (vehicle *Vehicle) Refuel(tank *Tank, quantity float64, date time.Time) error {
	if date.After(time.Now()) {
		return RefuelDateError
	}
	vehicleTank, err := NewVehicleTank(tank.id, vehicle.id, quantity, vehicle.odometer, date)
	if err != nil {
		return err
	}

	tank.quantity -= quantity
	vehicle.history = append(vehicle.history, vehicleTank)

	return nil
}
