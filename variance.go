package main

import (
	"fmt"
	"math"
)

func main() {
	data := []float64{50, 60, 70, 70, 100}

	total := float64(0)
	for _, d := range data {
		total += d
	}
	average := total / float64(len(data))

	fmt.Println("average: ", average)

	deviationTotal := float64(0)
	for index, d := range data {
		deviation := d - average
		fmt.Printf("deviation of index %d: %f\n", index, deviation)
		deviationTotal += math.Pow(deviation, 2)
	}

	variance := deviationTotal / float64(len(data))
	fmt.Println("variance: ", variance)

	standardVariance := math.Sqrt(variance)
	fmt.Println("standard variance: ", standardVariance)

	for index, d := range data {
		deviationScore := 10.0*(d-average)/standardVariance + 50
		fmt.Printf("deviation score of index %d: %f\n", index, deviationScore)
	}
}
