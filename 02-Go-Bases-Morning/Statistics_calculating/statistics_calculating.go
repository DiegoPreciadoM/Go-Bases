package main

import (
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Minimum(values ...int) float64 {
	min_value := values[0]

	for _, value := range values {
		if value < min_value {
			min_value = value
		}
	}
	return float64(min_value)
}

func Maximum(values ...int) float64 {
	max_value := values[0]

	for _, value := range values {
		if max_value < value {
			max_value = value
		}
	}
	return float64(max_value)
}

func Average(values ...int) float64 {
	values_sum := 0
	for _, value := range values {
		values_sum += value
	}

	return float64(values_sum) / float64(len(values))
}

func Choose_function(type_calculating string) func(...int) float64 {
	switch type_calculating {
	case minimum:
		return Minimum
	case average:
		return Average
	case maximum:
		return Maximum
	default:
		return nil
	}
}

func main() {
	minFunc := Choose_function(minimum)
	aveFunc := Choose_function(average)
	maxFunc := Choose_function(maximum)
	fmt.Println("1.- ", minFunc(2, 3, 3, 4, 10, 2, 4, 5))
	fmt.Println("2.- ", aveFunc(2, 3, 3, 4, 10, 2, 4, 5))
	fmt.Println("3.- ", maxFunc(2, 3, 3, 4, 10, 2, 4, 5))
}
