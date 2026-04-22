package phonenumber

import (
	"fmt"
)

type ErrInvalidPhoneNumberLength struct {
	NumberString string
}

func (e *ErrInvalidPhoneNumberLength) Error() string {
	return fmt.Sprintf("phone number has invalid length: %s (len: %d)", e.NumberString, len(e.NumberString))
}

type ErrInvalidAreaCodeStart struct {
	AreaCode string
}

func (e *ErrInvalidAreaCodeStart) Error() string {
	return fmt.Sprintf("area code must start from 2 to 9: %s", e.AreaCode)
}

type ErrInvalidExchangeCodeStart struct {
	ExchangeCode string
}

func (e *ErrInvalidExchangeCodeStart) Error() string {
	return fmt.Sprintf("exchange code must start from 2 to 9: %s", e.ExchangeCode)
}

type ErrInvalidCountryCode struct {
	CountryCode string
}

func (e *ErrInvalidCountryCode) Error() string {
	return fmt.Sprintf("country code must be 1: %s", e.CountryCode)
}

func parsePhoneNumber(phoneNumber string) (areaCode, exchangeCode, subscriberNumber string, err error) {
	numberString := ""
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			numberString += fmt.Sprintf("%d", (r - '0'))
		}
	}
	if len(numberString) < 10 {
		return "", "", "", &ErrInvalidPhoneNumberLength{numberString}
	}
	if len(numberString) > 10 {
		countryCode := numberString[:len(numberString)-10]
		if countryCode != "1" {
			return "", "", "", &ErrInvalidCountryCode{countryCode}
		}
	}
	numberString = numberString[len(numberString)-10:]
	areaCode = numberString[:3]
	if areaCode[0] < '2' {
		return "", "", "", &ErrInvalidAreaCodeStart{areaCode}
	}
	exchangeCode = numberString[3:6]
	if exchangeCode[0] < '2' {
		return "", "", "", &ErrInvalidExchangeCodeStart{areaCode}
	}
	subscriberNumber = numberString[6:10]
	err = nil
	return
}

func Number(phoneNumber string) (string, error) {
	areaCode, exchangeCode, subscriberNumber, err := parsePhoneNumber(phoneNumber)
	return areaCode + exchangeCode + subscriberNumber, err
}

func AreaCode(phoneNumber string) (string, error) {
	areaCode, _, _, err := parsePhoneNumber(phoneNumber)
	return areaCode, err
}

func Format(phoneNumber string) (string, error) {
	areaCode, exchangeCode, subscriberNumber, err := parsePhoneNumber(phoneNumber)
	format := "(%s) %s-%s"
	return fmt.Sprintf(format, areaCode, exchangeCode, subscriberNumber), err
}
