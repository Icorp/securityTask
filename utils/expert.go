package utils

import "log"

func GetExpertValue(n int) float64 {
	if n >= 21 && n <= 30 {
		return 0.01
	}
	if n >= 11 && n <= 20 {
		return 0.25
	}
	if n < 10 {
		return 0.5
	}
	if n == 0 {
		return 0.9
	}

	log.Printf("[ERROR] Cannot get Expert Value, input value:%d", n)
	return 1
}
