package enum

import "strings"

type VehicleType string

const (
	CAR    VehicleType = "CAR"
	TRUCK  VehicleType = "TRUCK"
	PICKUP VehicleType = "PICKUP"
)

func ValidVehicleType(value string) bool {
	switch VehicleType(value) {
	case CAR, TRUCK, PICKUP:
		return true
	}
	return false
}

func (tp VehicleType) ToString() string {
	switch tp {
	case CAR:
		return "CAR"
	case TRUCK:
		return "TRUCK"
	case PICKUP:
		return "PICKUP"
	default:
		return ""
	}
}

func GetVehicleTypes() string {
	types := []string{CAR.ToString(), TRUCK.ToString(), PICKUP.ToString()}
	result := strings.Join(types, ", ")
	return result
}
