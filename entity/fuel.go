package entity

import "github.com/google/uuid"

// Fuel entidade que abastece os tanques, cada uma possui a quantidade comprada em cada nota fiscal
type Fuel struct {
	ID            uuid.UUID
	Price         float64
	Quantity      float64
	InvoiceNumber string
}
