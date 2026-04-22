package phonenumber

import (
	"fmt"
)

type ErrInvalidPhoneNumberLength struct {
	digits []byte
}

func (e *ErrInvalidPhoneNumberLength) Error() string {
	return fmt.Sprintf("phone number has invalid length: %s (len: %d)", e.digits, len(e.digits))
}

type ErrInvalidAreaCodeStart struct {
	AreaCode []byte
}

func (e *ErrInvalidAreaCodeStart) Error() string {
	return fmt.Sprintf("area code must start from 2 to 9: %s", e.AreaCode)
}

type ErrInvalidExchangeCodeStart struct {
	ExchangeCode []byte
}

func (e *ErrInvalidExchangeCodeStart) Error() string {
	return fmt.Sprintf("exchange code must start from 2 to 9: %s", e.ExchangeCode)
}

type ErrInvalidCountryCode struct {
	CountryCode []byte
}

func (e *ErrInvalidCountryCode) Error() string {
	return fmt.Sprintf("country code must be 1: %s", e.CountryCode)
}

func parsePhoneNumber(phoneNumber string) ([]byte, []byte, []byte, error) {
	digits := make([]byte, 0, 10)
	for i := 0; i < len(phoneNumber); i++ {
		char := phoneNumber[i]
		if char >= '0' && char <= '9' {
			digits = append(digits, char)
		}
	}
	if len(digits) < 10 {
		return nil, nil, nil, &ErrInvalidPhoneNumberLength{digits}
	}
	if len(digits) == 11 && digits[0] != '1' {
		return nil, nil, nil, &ErrInvalidCountryCode{digits[:len(digits)-10]}
	}

	digits = digits[len(digits)-10:]
	areaCode := digits[:3]
	if areaCode[0] < '2' {
		return nil, nil, nil, &ErrInvalidAreaCodeStart{areaCode}
	}
	exchangeCode := digits[3:6]
	if exchangeCode[0] < '2' {
		return nil, nil, nil, &ErrInvalidExchangeCodeStart{exchangeCode}
	}
	subscriberNumber := digits[6:10]
	return areaCode, exchangeCode, subscriberNumber, nil
}

func Number(phoneNumber string) (string, error) {
	areaCode, exchangeCode, subscriberNumber, err := parsePhoneNumber(phoneNumber)
	return string(areaCode) + string(exchangeCode) + string(subscriberNumber), err
}

func AreaCode(phoneNumber string) (string, error) {
	areaCode, _, _, err := parsePhoneNumber(phoneNumber)
	return string(areaCode), err
}

func Format(phoneNumber string) (string, error) {
	areaCode, exchangeCode, subscriberNumber, err := parsePhoneNumber(phoneNumber)
	format := "(%s) %s-%s"
	return fmt.Sprintf(format, areaCode, exchangeCode, subscriberNumber), err
}
