package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen":    6,    
        "dozen":              12,
        "small_gross":        120,
        "gross":              144,
        "great_gross":        1728,
    }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; ok {
		bill[item] += units[unit]
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := bill[item]; !ok {
		return false // товара нет в счете
	}
	count, ok := units[unit]
	if !ok {
		return false // не нашли единицу измерения
	}
	if bill[item] < count {
		return false // не хватает товара для удаления
	}

	bill[item] -= count
	if bill[item] == 0 {
		delete(bill, item) // удаляем, если осталось 0
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, ok := bill[item]; ok {
		return bill[item], true
	}
	return 0, false	
}
