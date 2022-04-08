package cars

const batchCost uint = 95_000
const singleCost uint = 10_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) (workingCars float64) {
	workingCars = (float64(productionRate) / 100) * successRate
    return
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) (workingCars int) {
    workingCars = int(float64(productionRate / 60) * (successRate / 100))
    return
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) (totalCost uint) {
    carsCountUint := uint(carsCount)
    batches := carsCountUint / 10
    rest := carsCountUint % 10

    totalCost = batches * batchCost + rest * singleCost
    return
}
