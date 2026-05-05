package proteintranslation

import (
	"errors"
	"strings"
)

const (
	methionine    = "Methionine"
	phenylalanine = "Phenylalanine"
	leucine       = "Leucine"
	serine        = "Serine"
	tyrosine      = "Tyrosine"
	cysteine      = "Cysteine"
	tryptophan    = "Tryptophan"
	stop          = "STOP"
)

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

var codonToAminoAcid = map[string]string{
	"aug": methionine,
	"uuu": phenylalanine, "uuc": phenylalanine,
	"uua": leucine, "uug": leucine,
	"ucu": serine, "ucc": serine, "uca": serine, "ucg": serine,
	"uau": tyrosine, "uac": tyrosine,
	"ugu": cysteine, "ugc": cysteine,
	"ugg": tryptophan,
	"uaa": stop, "uag": stop, "uga": stop,
}

func FromRNA(rna string) ([]string, error) {
	var aas []string

	remaining := rna
	for len(remaining) >= 3 {
		codon := remaining[:3]
		aa, err := FromCodon(codon)
		switch err {
		case nil:
			aas = append(aas, aa)
		case ErrInvalidBase:
			return nil, err
		case ErrStop:
			return aas, nil
		}
		remaining = remaining[3:]
	}
	if len(remaining) > 0 {
		return nil, ErrInvalidBase
	}
	return aas, nil
}

func FromCodon(codon string) (string, error) {
	aminoAcid, found := codonToAminoAcid[strings.ToLower(codon)]
	if !found {
		return "", ErrInvalidBase
	}
	if aminoAcid == stop {
		return "", ErrStop
	}
	return aminoAcid, nil
}
