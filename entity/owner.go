package entity

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
	"strings"
	"time"
)

var (
	OwnerCnpjError           = errors.New("invalid cnpj")
	OwnerCoorporateNameError = errors.New("invalid coorporate name")
	OwnerEmailError          = errors.New("invalid email")
)

type Owner struct {
	Base
	cnpj           string
	coorporateName string
	email          string
}

func NewOwner(cnpj, coorporateName, email string) (Owner, error) {
	err := validateOwner(cnpj, coorporateName, email)
	if err != nil {
		return Owner{}, err
	}
	return Owner{
		Base: Base{
			id:        uuid.New(),
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		cnpj:           cnpj,
		coorporateName: coorporateName,
		email:          email,
	}, nil
}

func validateOwner(cnpj, coorporateName, email string) error {
	if cnpj == "" || !validateCnpj(cnpj) {
		return OwnerCnpjError
	}
	if coorporateName == "" {
		return OwnerCoorporateNameError
	}
	if email == "" || !strings.ContainsAny(email, "@") {
		return OwnerEmailError
	}
	return nil
}

func validateCnpj(cnpj string) bool {
	re := regexp.MustCompile(`[^\d]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	if len(cnpj) != 14 {
		return false
	}

	sum := 0
	for i := 0; i < 12; i++ {
		sum += int(cnpj[i]-'0') * (i + 2)
	}
	remainder := sum % 11
	digit1 := 11 - remainder
	if digit1 >= 10 {
		digit1 = 0
	}
	if digit1 != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		sum += int(cnpj[i]-'0') * (i + 2)
	}
	remainder = sum % 11
	digit2 := 11 - remainder
	if digit2 >= 10 {
		digit2 = 0
	}
	if digit2 != int(cnpj[13]-'0') {
		return false
	}

	return true
}
