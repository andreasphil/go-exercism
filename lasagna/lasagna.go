package lasagna

// How many minutes the lasagna should be in the oven
const OvenTime = 40

// Returns how many minutes the lasagna still has to be in the oven
func RemainingOvenTime(minutesInOven int) int {
	return OvenTime - minutesInOven
}

// Returns how many minutes you spent preparing the lasagna, assuming each
// layer takes you 2 minutes to prepare
func PreparationTime(layers int) int {
	return layers * 2
}

// Return how many minutes in total you've worked on cooking the lasagna,
// which is the sum of the preparation time in minutes, and the time in
// minutes the lasagna has spent in the oven at the moment
func ElapsedTime(layers, minutesInOven int) int {
	return PreparationTime(layers) + minutesInOven
}
