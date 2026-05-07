package gross

var units = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitAmount, unitExists := units[unit]
	if !unitExists {
		return false
	}
	bill[item] += unitAmount
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitAmount, unitExists := units[unit]
	if !unitExists {
		return false
	}
	_, itemInBill := bill[item]
	if !itemInBill {
		return false
	}
	if bill[item]-unitAmount < 0 {
		return false
	}
	bill[item] -= unitAmount
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemAmount, itemInBill := bill[item]
	if !itemInBill {
		return 0, false
	}
	return itemAmount, true
}
