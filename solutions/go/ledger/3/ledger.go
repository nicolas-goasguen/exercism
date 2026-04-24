package ledger

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	curUSD = "USD"
	curEUR = "EUR"

	signEUR = "€"
	signUSD = "$"

	locUS = "en-US"
	locNL = "nl-NL"

	thousandSepUS = ","
	thousandSepNL = "."

	decimalSepUS = "."
	decimalSepNL = ","

	dateFmtShort = "2006-01-02"
	dateFmtUS    = "01/02/2006"
	dateFmtNL    = "02-01-2006"
)

var curSigns = map[string]string{
	curUSD: signUSD,
	curEUR: signEUR,
}

var locDateFmt = map[string]string{
	locUS: dateFmtUS,
	locNL: dateFmtNL,
}

var locThousandSep = map[string]string{
	locUS: thousandSepUS,
	locNL: thousandSepNL,
}

var locDecimalSep = map[string]string{
	locUS: decimalSepUS,
	locNL: decimalSepNL,
}

var currencies = []string{curUSD, curEUR}
var locales = []string{locUS, locNL}

var (
	ErrInvalidCurrency = errors.New("currency is invalid")
	ErrInvalidLocale   = errors.New("locale is invalid")
	ErrInvalidDate     = errors.New("date is invalid")
	ErrInvalidChange   = errors.New("change is invalid")
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func (e Entry) Compare(other Entry) int {
	if e.Date != other.Date {
		return strings.Compare(e.Date, other.Date)
	}
	if e.Description != other.Description {
		return strings.Compare(e.Description, other.Description)
	}
	return cmp.Compare(e.Change, other.Change)
}

func formatHeader(locale string) (string, error) {
	var h strings.Builder
	var date string
	var description string
	var change string

	switch locale {
	case locNL:
		date = "Datum"
		description = "Omschrijving"
		change = "Verandering"
	case locUS:
		date = "Date"
		description = "Description"
		change = "Change"
	default:
		return "", ErrInvalidLocale
	}

	h.WriteString(fmt.Sprintf("%-10s |", date))
	h.WriteString(fmt.Sprintf(" %-25s |", description))
	h.WriteString(fmt.Sprintf(" %-13s\n", change))
	return h.String(), nil
}

func formatDate(locale, date string) (string, error) {
	t, err := time.Parse(dateFmtShort, date)
	if err != nil {
		return "", err
	}

	df, ok := locDateFmt[locale]
	if !ok {
		return "", ErrInvalidLocale
	}

	return t.Format(df), nil
}

func formatDescription(desc string) string {
	var d string
	r := []rune(desc)
	if len(r) <= 25 {
		d = fmt.Sprintf("%-25s", desc)
	} else {
		d = fmt.Sprintf("%s...", string(r[:22]))
	}
	return d
}

// units must be non-negative
func formatChangeUnits(units int, thousandSep string) string {
	var unitsStrBuilder strings.Builder
	unitsStr := strconv.Itoa(units)
	for i, digit := range unitsStr {
		if i > 0 && (len(unitsStr)-i)%3 == 0 {
			unitsStrBuilder.WriteString(thousandSep)
		}
		unitsStrBuilder.WriteRune(digit)
	}
	return unitsStrBuilder.String()
}

func formatChangeCents(cents int) string {
	return fmt.Sprintf("%02d", cents)
}

func formatChange(currency string, locale string, change int) (string, error) {
	var negative bool
	absChange := change
	if absChange < 0 {
		absChange *= -1
		negative = true
	}

	units := absChange / 100
	cents := absChange % 100

	curSign, ok := curSigns[currency]
	if !ok {
		return "", ErrInvalidCurrency
	}

	thousandSep, ok := locThousandSep[locale]
	if !ok {
		return "", ErrInvalidLocale
	}
	decimalSep, ok := locDecimalSep[locale]
	if !ok {
		return "", ErrInvalidLocale
	}

	fmtUnits := formatChangeUnits(units, thousandSep)
	fmtCents := formatChangeCents(cents)

	var format string
	var fmtChange string
	switch locale {
	case locUS:
		if negative {
			format = "(%s%s%s%s)"
		} else {
			format = " %s%s%s%s "
		}
	case locNL:
		if negative {
			format = "%s -%s%s%s "
		} else {
			format = "%s %s%s%s "
		}
	default:
		return "", ErrInvalidLocale
	}
	fmtChange = fmt.Sprintf(format, curSign, fmtUnits, decimalSep, fmtCents)
	return fmtChange, nil
}

func formatEntry(currency string, locale string, entry *Entry) (string, error) {
	var errs []error

	fmtDate, errDate := formatDate(locale, entry.Date)
	if errDate != nil {
		errs = append(errs, errDate)
	}

	fmtDesc := formatDescription(entry.Description)

	fmtChange, errChange := formatChange(currency, locale, entry.Change)
	if errChange != nil {
		errs = append(errs, errChange)
	}

	if len(errs) != 0 {
		return "", errors.Join(errs...)
	}

	fmtEntry := fmt.Sprintf("%10s | %s | %13s\n", fmtDate, fmtDesc, fmtChange)
	return fmtEntry, nil
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if !slices.Contains(currencies, currency) {
		return "", ErrInvalidCurrency
	}
	if !slices.Contains(locales, locale) {
		return "", ErrInvalidLocale
	}

	entriesCopy := slices.Clone(entries)
	slices.SortFunc(entriesCopy, Entry.Compare)

	fmtHeader, err := formatHeader(locale)
	if err != nil {
		return "", err
	}

	var ledgerStrBuilder strings.Builder
	ledgerStrBuilder.WriteString(fmtHeader)

	for _, entry := range entriesCopy {
		fmtEntry, err := formatEntry(currency, locale, &entry)
		if err != nil {
			return "", err
		}
		ledgerStrBuilder.WriteString(fmtEntry)
	}

	return ledgerStrBuilder.String(), nil
}
