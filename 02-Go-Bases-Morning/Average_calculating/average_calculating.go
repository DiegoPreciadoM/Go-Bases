package main

import "fmt"

func average_calculating(values ...int) float64 {

	values_sum := 0
	for _, value := range values {
		if value < 0 {
			return 0.0
		}
		values_sum += value
	}

	average := float64(values_sum) / float64(len(values))

	return float64(average)
}

func main() {
	fmt.Printf("\nEl promedio es: %f\n", average_calculating(9, 8, 9, 7, 8, 9, 8, 9))
}
