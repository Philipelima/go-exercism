package cars

const HOUR_MINUTES int = 60

const TEN_CARS_GROUP_COST int = 95000
const CAR_UNIT_COST int = 10000


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

    avg  :=  (successRate / 100) * float64(productionRate)
    return float64(avg)
}


// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

    carsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	carsPerMinutes := int(carsPerHour) / HOUR_MINUTES

    return carsPerMinutes
}


// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

    groups    := carsCount / 10 
    remaining := carsCount % 10
  
    cost := (groups * TEN_CARS_GROUP_COST) + (remaining * CAR_UNIT_COST)
	return uint(cost)
}