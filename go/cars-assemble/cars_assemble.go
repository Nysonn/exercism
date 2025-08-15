package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	numberOfSuccessfullCarsProducedPerHour := float64(productionRate) * float64(successRate/100)
	return numberOfSuccessfullCarsProducedPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	numberOfSuccessfullCarsProducedPerHour := float64(productionRate) * float64(successRate/100)
	numberOfSuccessfullCarsProducedPerMinute := int(numberOfSuccessfullCarsProducedPerHour / 60)
	return numberOfSuccessfullCarsProducedPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	individualCars := carsCount % 10

	costOfProducingCars := uint(groupsOfTen*95000 + individualCars*10000)
	return costOfProducingCars
}
