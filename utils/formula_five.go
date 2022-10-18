package utils

// Велечина вероятного ущерба
func CalculateL(a []int, b []int) int {
	bossPrices := 0
	expertPrices := 0

	for _, i2 := range a {
		bossPrices += i2
	}

	for _, i2 := range b {
		expertPrices += i2
	}

	return bossPrices + expertPrices
}
