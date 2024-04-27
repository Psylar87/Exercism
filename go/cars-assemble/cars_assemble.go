package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var perHourSuccess = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHourSuccess) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var perCar = 10000
	var groupCost = 95000
	var groups = carsCount / 10
	var remaining = carsCount % 10
	return uint(groups) * uint(groupCost) + uint(remaining) * uint(perCar)
}
