package isbnverifier

const separator = '-'

func IsValidISBN(isbn string) bool {
	sum := 0
	count := 0
	for i := range len(isbn) {
		r := isbn[i]
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digit := int(r - '0')
			sum += digit * (10 - count)
			count++
		case separator:
			continue
		case 'X':
			if count != 9 || i != len(isbn)-1 {
				return false
			}
			sum += 10
			count++
		default:
			return false
		}
	}
	return count == 10 && sum%11 == 0
}
