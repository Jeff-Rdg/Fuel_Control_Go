package entity

// Fuel entidade que abastece os tanques, cada uma possui a quantidade comprada em cada nota fiscal
type Fuel struct {
	Base
	Price         float64
	Quantity      float64
	InvoiceNumber string
}
