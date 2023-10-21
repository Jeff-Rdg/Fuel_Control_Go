package test

import (
	"FuelControl/entity"
	"errors"
	"testing"
)

func TestFuel_NewFuel(t *testing.T) {
	type testCase struct {
		testName      string
		price         float64
		quantity      float64
		invoiceNumber string
		expectedErr   error
	}

	testCases := []testCase{
		{
			testName:      "Negative Price",
			price:         -5.0,
			quantity:      5.0,
			invoiceNumber: "123456",
			expectedErr:   entity.FuelPriceError,
		},
		{
			testName:      "Zero Price",
			price:         0.0,
			quantity:      5.0,
			invoiceNumber: "123456",
			expectedErr:   entity.FuelPriceError,
		},
		{
			testName:      "Negative quantity",
			price:         5.0,
			quantity:      -5.0,
			invoiceNumber: "123456",
			expectedErr:   entity.FuelQuantityError,
		},
		{
			testName:      "Zero quantity",
			price:         5.0,
			quantity:      0.0,
			invoiceNumber: "123456",
			expectedErr:   entity.FuelQuantityError,
		},
		{
			testName:      "invoiceNumber with five characters",
			price:         5.0,
			quantity:      1.0,
			invoiceNumber: "12345",
			expectedErr:   entity.FuelInvoiceNumberError,
		},
		{
			testName:      "invoiceNumber with seven characters",
			price:         5.0,
			quantity:      1.0,
			invoiceNumber: "1234567",
			expectedErr:   entity.FuelInvoiceNumberError,
		},
		{
			testName:      "invoiceNumber with zero characters",
			price:         5.0,
			quantity:      1.0,
			invoiceNumber: "",
			expectedErr:   entity.FuelInvoiceNumberError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewFuel(tc.price, tc.quantity, tc.invoiceNumber)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
