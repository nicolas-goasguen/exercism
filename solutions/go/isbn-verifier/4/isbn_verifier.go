package isbnverifier

const separator = '-'

func IsValidISBN(isbn string) bool {
	sum := 0
	count := 0
	for i := range len(isbn) {
		if count+len(isbn)-i < 10 {
			return false
		}
		r := isbn[i]
		if r >= '0' && r <= '9' {
			digit := int(r - '0')
			sum += digit * (10 - count)
			count++
		} else if r == separator {
			continue
		} else if r == 'X' && count == 9 {
			sum += 10
			count++
		} else {
			return false
		}
	}
	return count == 10 && sum%11 == 0
}
