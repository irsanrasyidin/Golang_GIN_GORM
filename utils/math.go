package utils

import (
	"math"
)

func Average(data []int) int {
	var sum int
	for _, value := range data {
		sum += value
	}
	return sum / len(data)
}

func AverageDuration(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	return math.Round(sum/float64(len(data))*1000) / 1000
}

func StandardDeviation(data []float64) float64 {
	// Calculate the mean
	var mean float64
	for _, value := range data {
		mean += value
	}
	mean /= float64(len(data))

	// Calculate the sum of squares of differences from the mean
	var sumOfSquares float64
	for _, value := range data {
		diff := value - mean
		sumOfSquares += diff * diff
	}

	// Calculate the variance
	variance := sumOfSquares / float64(len(data)-1)

	// Calculate the standard deviation (square root of the variance)
	standardDeviation := math.Sqrt(variance)

	return math.Round(standardDeviation*1000) / 1000
}

func Filter(data []float64, outliner int) ([]float64, int, float64) {
	// Temukan nilai terbesar dan indeksnya
	max := data[0]
	maxIndex := 0
	for i, val := range data {
		if val > max {
			max = val
			maxIndex = i
		}
	}

	// Hapus nilai terbesar dari array
	data = append(data[:maxIndex], data[maxIndex+1:]...)
	outliner++
	deletedData := max
	return data, outliner, deletedData
}
