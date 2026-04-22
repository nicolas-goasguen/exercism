package isbnverifier

var areSeparators13 = [13]bool{1: true, 5: true, 11: true}

const separator = '-'

func IsValidISBN(isbn string) bool {
	if len(isbn) != 13 && len(isbn) != 10 {
		return false
	}

	var withSep bool
	if len(isbn) == 13 {
		withSep = true
	}

	sum := 0
	multi := 10
	for i := range len(isbn) {
		r := isbn[i]
		if r >= '0' && r <= '9' {
			digit := int(r - '0')
			sum += digit * multi
			multi--
		} else if r == separator {
			if withSep && areSeparators13[i] {
				continue
			} else {
				return false
			}
		} else if r == 'X' && i == len(isbn)-1 {
			sum += 10
		} else {
			return false
		}
	}
	return sum%11 == 0
}
