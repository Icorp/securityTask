package main

import (
	"github.com/Icorp/securityTask/utils"
	"log"
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
	log.Printf("Угрозы по доступности и целостности можно посчитать по формуле(2): %.2f", probabilityResult)

	// Результат вычисления P2(A|B)
	result := calculateAFromB(selectedActive)
	log.Printf("Условная вероятность: %.2f", result)

	// Результат формула зависимых событий Р2(А1B)
	result2 := dependenceEvent(selectedActive, result)
	log.Printf("Тогда, по формуле зависимых событий: %.2f", result2)

	a := append(data2, result2)

	// Результат Р2(А1)
	probabilityResult2 := utils.CalculateProbability(a)
	log.Printf("Вероятность угрозы по конфиденциальности: %.2f", probabilityResult2)

	// Получаем экспертную оценку.
	expertValue := utils.GetExpertValue(numOfFix)
	log.Printf("По экспертным оценкам, вероятность угрозы (Рz) составляет: %.2f", expertValue)

	// Вероятность угрозы P общ
	probabilityTotal := utils.CalculateProbability([]float64{probabilityResult, probabilityResult2, expertValue})
	log.Printf("C вероятностью %.2f с активом А1 произойдет хотя бы одно неблагоприятное событие из списка всех актуальных угроз.", probabilityTotal)

	// Величина вероятного ущерба
	loss := utils.CalculateL(bossPrices, expertPrices)
	log.Printf("Величина вероятного ущерба по формуле(5): %d", loss)

	// Риск
	risk := utils.CalculateR(probabilityTotal, loss)
	log.Printf("Риск вычисленный по формуле: %.2f", risk)
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
