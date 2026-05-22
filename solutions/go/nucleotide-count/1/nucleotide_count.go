package nucleotidecount

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
// type Histogram ...

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
// type DNA ...

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.

var (
	ErrInvalidNucleotide = errors.New("invalid nucleotide")
)

type Histogram map[rune]int

func NewHistogram() Histogram {
	var h Histogram = make(Histogram)
	for _, r := range []rune{'A', 'C', 'G', 'T'} {
		h[r] = 0
	}
	return h
}

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := NewHistogram()
	for _, r := range d {
		if _, ok := h[r]; !ok {
			return nil, ErrInvalidNucleotide
		}
		h[r]++
	}
	return h, nil
}
