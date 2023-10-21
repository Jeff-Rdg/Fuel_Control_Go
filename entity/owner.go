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
	multiplier := 5
	for i := 0; i < 12; i++ {
		digit := int(cnpj[i] - '0')
		sum += digit * multiplier
		multiplier--
		if multiplier < 2 {
			multiplier = 9
		}
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}

	if digit1 != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	multiplier = 6
	for i := 0; i < 13; i++ {
		digit := int(cnpj[i] - '0')
		sum += digit * multiplier
		multiplier--
		if multiplier < 2 {
			multiplier = 9
		}
	}
	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}

	if digit2 != int(cnpj[13]-'0') {
		return false
	}

	return true
}
