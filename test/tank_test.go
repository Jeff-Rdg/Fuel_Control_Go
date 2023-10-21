package test

import (
	"FuelControl/entity"
	"errors"
	"testing"
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
