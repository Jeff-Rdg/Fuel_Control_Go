package test

import (
	"FuelControl/entity"
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestFuelTank_NewFuelTank(t *testing.T) {
	type testCase struct {
		testName    string
		fromFuel    uuid.UUID
		toTank      uuid.UUID
		quantity    float64
		supplyDate  time.Time
		expectedErr error
	}

	testCases := []testCase{
		{
			testName:    "Negative quantity",
			fromFuel:    uuid.New(),
			toTank:      uuid.New(),
			quantity:    -5,
			supplyDate:  time.Now(),
			expectedErr: entity.FuelTankQuantityError,
		},
		{
			testName:    "Valid quantity",
			fromFuel:    uuid.New(),
			toTank:      uuid.New(),
			quantity:    5,
			supplyDate:  time.Now(),
			expectedErr: nil,
		},
		{
			testName:    "invalid date",
			fromFuel:    uuid.New(),
			toTank:      uuid.New(),
			quantity:    5,
			supplyDate:  time.Now().AddDate(0, 0, 1),
			expectedErr: entity.FuelTankDateError,
		},
		{
			testName:    "valid date",
			fromFuel:    uuid.New(),
			toTank:      uuid.New(),
			quantity:    5,
			supplyDate:  time.Now().AddDate(0, 0, -1),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewFuelTank(tc.fromFuel, tc.toTank, tc.quantity, tc.supplyDate)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
