package cars


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateCost(carsCount int) uint {
	return uint((int(carsCount/10) * 95000) + ((carsCount - (int(carsCount/10) * 10)) * 10000))
}
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    workingCarsPerHour := float64(productionRate) * successRate / 100
    workingCarsPerMinute := workingCarsPerHour / 60
    return int(workingCarsPerMinute)  // Простое усечение
}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}