package enum

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
