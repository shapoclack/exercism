package purchase

func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func ChooseVehicle(option1, option2 string) string {
	var chosen string
	
	// Сравниваем строки в лексикографическом порядке
	if option1 < option2 {
		chosen = option1
	} else {
		chosen = option2
	}
	
	result := chosen + " is clearly the better choice."
	return result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var result float64
	if age < 3 {
		result = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		result = originalPrice * 0.7
	} else if age >= 10 {
		result = originalPrice/100 * 50
	}
	return result
}

func main() {
	answer1 := CalculateResellPrice(1000, 10)
	println(answer1)
	answer2 := NeedsLicense("car")
	println(answer2)
	answer3 := ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
	println(answer3)
}
