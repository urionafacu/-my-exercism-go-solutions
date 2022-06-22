package lasagna

const (
	OvenTime = 40
)

func RemainingOvenTime(actual int) int {
	if actual > OvenTime {
		return 0
	} else if actual < 0 {
		return 40
	}
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	if numberOfLayers <= 0 {
		return 0
	}
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
