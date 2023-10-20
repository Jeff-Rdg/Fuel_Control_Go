package entity

import (
	"github.com/google/uuid"
	"time"
)

type VehicleTank struct {
	FromTank      uuid.UUID
	ToVehicle     uuid.UUID
	Quantity      float64
	Odometer      float64
	ProvisionDate time.Time
}
