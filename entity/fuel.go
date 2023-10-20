package entity

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
	"time"
)

var (
	FuelPriceError         = errors.New("incorrect value of price")
	FuelQuantityError      = errors.New("incorrect value of quantity")
	FuelInvoiceNumberError = errors.New("incorrect Invoice Number")
)

// Fuel entidade que abastece os tanques, cada uma possui a quantidade comprada em cada nota fiscal
type Fuel struct {
	Base
	price         float64
	quantity      float64
	invoiceNumber string
}

func NewFuel(price float64, quantity float64, invoiceNumber string) (Fuel, error) {
	err := validateNewFuel(price, quantity, invoiceNumber)
	if err != nil {
		return Fuel{}, err
	}

	return Fuel{
		Base: Base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		price:         price,
		quantity:      quantity,
		invoiceNumber: invoiceNumber,
	}, nil
}

func (fuel Fuel) TotalPrice() float64 {
	return fuel.price * fuel.quantity
}

func validateNewFuel(price, quantity float64, invoiceNumber string) error {
	if price <= 0 {
		return FuelPriceError
	}
	if quantity <= 0 {
		return FuelQuantityError
	}

	regex := regexp.MustCompile("^[0-9]+$")
	if invoiceNumber == "" || !regex.MatchString(invoiceNumber) || len(invoiceNumber) != 6 {
		return FuelInvoiceNumberError
	}

	return nil
}
