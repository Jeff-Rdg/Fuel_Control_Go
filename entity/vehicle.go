package entity

import (
	"FuelControl/entity/enum"
)

type Vehicle struct {
	Base
	plate       string
	odometer    int64
	vehicleType enum.VehicleType
	Owner
}
