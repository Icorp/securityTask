package main

import (
	"fmt"
	"github.com/Icorp/securityTask/utils"
)

// P - количественная оценка возможности наступления хотя бы одного события (из всех угроз на актив An);
// A - Активы
// Lобщ(An) - общая стоимость потери актива An;
// R(An) - количественное выражение риска на актив An;
// C(An) - количественное выражение ущерба в деньгах от реализации угрозы на актив An;
// F(An) - затраты на восстановление актива

const factor = 0.7
const humanFactor = 0.75
const numOfFix = 15 // Количество выполненных мер

func main() {
	data := []float64{0.1, 0.2, 0.05, 0.05, 0.1, 0.05, 0.07, 0.2} // Угроза доступности + Угрозы целостности
	data2 := []float64{0.15, 0.1, 0.1}                            // Угрозы конфиденциальности
	bossPrices := []int{500000, 1000000}                          // Оценка руководства
	expertPrices := []int{300000, 100000}                         // Оценка эксперта
	selectedActive := 0.15                                        // Выбранный актив

	// Результат вычисления p(a)
	probabilityResult := utils.CalculateProbability(data)

	// Результат вычисления P2(A|B)
	result := calculateAFromB(selectedActive)

	// Результат формула зависимых событий Р2(А1B)
	result2 := dependenceEvent(selectedActive, result)

	a := append(data2, result2)

	// Результат Р2(А1)
	probabilityResult2 := utils.CalculateProbability(a)

	// Получаем экспертную оценку.
	expertValue := utils.GetExpertValue(numOfFix)

	// Вероятность угрозы P общ
	pTotal := utils.CalculateProbability([]float64{probabilityResult, probabilityResult2, expertValue})

	// Величина вероятного ущерба
	loss := calculateL(bossPrices, expertPrices)

	// Риск
	risk := calculateR(pTotal, loss)
	fmt.Println(risk)
}

// Велечина вероятного ущерба
func calculateL(a []int, b []int) int {
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

// Риск(количественная величина)
func calculateR(a float64, b int) float64 {
	return a * float64(b)
}

// по формуле зависимых событий
func dependenceEvent(a float64, b float64) float64 {
	return a * b
}

// Результат вычисления P2(A|B)
func calculateAFromB(b float64) float64 {
	a := []float64{b, factor, humanFactor}
	return utils.CalculateProbability(a)
}
