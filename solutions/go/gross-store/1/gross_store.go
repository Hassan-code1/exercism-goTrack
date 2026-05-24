package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	allUnits := map[string]int{}
    allUnits["quarter_of_a_dozen"] = 3
    allUnits["half_of_a_dozen"] = 6
    allUnits["dozen"] = 12
    allUnits["small_gross"] = 120
    allUnits["gross"] = 144
    allUnits["great_gross"] = 1728
    return allUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _ , exist := units[unit]; !exist {
        return false
    }
    if _, exist := bill[item]; exist{
        bill[item] += units[unit]
    }else{
        bill[item] = units[unit]
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, exist := bill[item]; !exist{
        return false
    }
    if _,exist := units[unit]; !exist{
        return false
    }
    if bill[item] - units[unit] < 0 {
        return false
    }
    if bill[item] - units[unit] == 0 {
        delete(bill, item)
        return true
    }
    bill[item] -= units[unit]
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, exist := bill[item]; !exist{
        return 0, false
    }
    return bill[item], true
}
