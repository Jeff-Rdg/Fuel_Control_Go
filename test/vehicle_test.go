package test

import (
	"FuelControl/entity"
	"errors"
	"testing"
)

func TestVehicle_NewVehicle(t *testing.T) {
	type testCase struct {
		testName    string
		plate       string
		odometer    int
		vehicleType string
		owner       entity.Owner
		expectedErr error
	}

	validOwner, _ := entity.NewOwner("40.337.128/0001-60", "Test Owner LTDA", "test@gmail.com")

	testCases := []testCase{
		{
			testName:    "Invalid plate",
			plate:       "",
			odometer:    1,
			vehicleType: "TRUCK",
			owner:       validOwner,
			expectedErr: entity.VehiclePlateError,
		},
		{
			testName:    "Valid plate",
			plate:       "MVJ7605",
			odometer:    1,
			vehicleType: "TRUCK",
			owner:       validOwner,
			expectedErr: nil,
		},
		{
			testName:    "Invalid odometer",
			plate:       "MVJ7605",
			odometer:    0,
			vehicleType: "TRUCK",
			owner:       validOwner,
			expectedErr: entity.VehicleOdometerError,
		},
		{
			testName:    "Valid odometer",
			plate:       "MVJ7605",
			odometer:    150,
			vehicleType: "TRUCK",
			owner:       validOwner,
			expectedErr: nil,
		},
		{
			testName:    "Invalid vehicle Type",
			plate:       "MVJ7605",
			odometer:    150,
			vehicleType: "TRUCKA",
			owner:       validOwner,
			expectedErr: entity.VehicleTypeError,
		},
		{
			testName:    "valid vehicle Type",
			plate:       "MVJ7605",
			odometer:    150,
			vehicleType: "CAR",
			owner:       validOwner,
			expectedErr: nil,
		},
		{
			testName:    "Invalid owner",
			plate:       "MVJ7605",
			odometer:    150,
			vehicleType: "CAR",
			expectedErr: entity.VehicleOwnerError,
		},
		{
			testName:    "Valid owner",
			plate:       "MVJ7605",
			odometer:    150,
			vehicleType: "CAR",
			owner:       validOwner,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewVehicle(tc.plate, tc.odometer, tc.vehicleType, tc.owner)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
