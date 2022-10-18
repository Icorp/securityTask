package utils

// Формула 2
func CalculateProbability(a []float64) float64 {
	partOneValue := calculatePartOne(a)
	partTwoValue := calculatePartTwo(a)
	partThreeValue := calculatePartThree(a)
	return partOneValue + partTwoValue + partThreeValue
}

func calculatePartOne(listA []float64) float64 {
	result := 0.0
	for _, i2 := range listA {
		result += i2
	}

	return result
}
func calculatePartTwo(listA []float64) float64 {
	result := 0.0
	for k := range listA {
		cash := 0.0
		for j := range listA {
			if k != j {
				cash += listA[j]
			}
		}
		result -= cash * listA[k]
	}

	return result
}
func calculatePartThree(listA []float64) float64 {
	result := 1.0
	for _, value := range listA {
		result *= value
	}

	return result
}
