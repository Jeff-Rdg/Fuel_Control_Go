package test

import (
	"FuelControl/entity"
	"errors"
	"testing"
	"time"
)

func TestTank_NewTank(t *testing.T) {
	type testCase struct {
		testName    string
		capacity    float64
		expectedErr error
	}

	testCases := []testCase{
		{
			testName:    "Negative capacity",
			capacity:    -5.0,
			expectedErr: entity.TankCapacityError,
		},
		{
			testName:    "Zero capacity",
			capacity:    0.0,
			expectedErr: entity.TankCapacityError,
		},
		{
			testName:    "capacity greater than 2000",
			capacity:    2001,
			expectedErr: entity.TankCapacityError,
		},
		{
			testName:    "capacity equals 2000",
			capacity:    2000,
			expectedErr: nil,
		},
		{
			testName:    "capacity valid",
			capacity:    1500.84,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewTank(tc.capacity)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestTank_Refuel(t *testing.T) {
	type testCase struct {
		testName    string
		tank        entity.Tank
		fuel        entity.Fuel
		date        time.Time
		expectedErr error
	}

	fuel, _ := entity.NewFuel(5, 100, "123456")
	tank, _ := entity.NewTank(1000)

	testCases := []testCase{
		{
			testName:    "Valid Refuel",
			tank:        tank,
			fuel:        fuel,
			date:        time.Now(),
			expectedErr: nil,
		},
		{
			testName:    "Invalid date to refuel",
			tank:        tank,
			fuel:        fuel,
			date:        time.Now().AddDate(0, 0, 1),
			expectedErr: entity.TankRefuelDateError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := tank.Refuel(tc.fuel, tc.date)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

	tank.SetCapacity(100)
	testCases = []testCase{
		{
			testName:    "invalid refuel quantity fuel greater than tank capacity",
			tank:        tank,
			fuel:        fuel,
			date:        time.Now(),
			expectedErr: entity.TankRefuelQuantityError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			err := tank.Refuel(tc.fuel, tc.date)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
