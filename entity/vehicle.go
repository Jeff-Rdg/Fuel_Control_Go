package entity

import (
	"FuelControl/entity/enum"
)

type Vehicle struct {
	Base
	Plate    string
	Odometer int64
	Type     enum.VehicleType
	Owner
}
