package test

import (
	"FuelControl/entity"
	"errors"
	"testing"
)

func TestOwner_NewOwner(t *testing.T) {
	type testCase struct {
		testName       string
		cnpj           string
		coorporateName string
		email          string
		expectedErr    error
	}

	testCases := []testCase{
		{
			testName:       "Invalid cnpj",
			cnpj:           "12345678901234",
			coorporateName: "Test Owner LTDA",
			email:          "test@gmail.com",
			expectedErr:    entity.OwnerCnpjError,
		},
		{
			testName:       "Valid cnpj",
			cnpj:           "40.337.128/0001-60",
			coorporateName: "Test Owner LTDA",
			email:          "test@gmail.com",
			expectedErr:    nil,
		},
		{
			testName:       "Invalid coorporateName",
			cnpj:           "71.430.236/0001-06",
			coorporateName: "",
			email:          "test@gmail.com",
			expectedErr:    entity.OwnerCoorporateNameError,
		},
		{
			testName:       "Valid coorporateName",
			cnpj:           "71.430.236/0001-06",
			coorporateName: "Test Owner LTDA",
			email:          "test@gmail.com",
			expectedErr:    nil,
		},
		{
			testName:       "Invalid email without @",
			cnpj:           "71.430.236/0001-06",
			coorporateName: "Test Owner LTDA",
			email:          "test.gmail.com",
			expectedErr:    entity.OwnerEmailError,
		},
		{
			testName:       "Invalid email",
			cnpj:           "71.430.236/0001-06",
			coorporateName: "Test Owner LTDA",
			email:          "",
			expectedErr:    entity.OwnerEmailError,
		},
		{
			testName:       "Valid email",
			cnpj:           "71.430.236/0001-06",
			coorporateName: "Test Owner LTDA",
			email:          "test@gmail.com",
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := entity.NewOwner(tc.cnpj, tc.coorporateName, tc.email)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
