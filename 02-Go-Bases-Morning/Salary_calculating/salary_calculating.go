package main

import "fmt"

func salary_calculating(mins int, category string) float64 {

	hours := mins / 60

	switch category {
	case "A":
		return float64(float64(hours*3000) + (float64(hours*3000) * 0.5))
	case "B":
		return float64(float64(hours*1500) + (float64(hours*1500) * 0.2))
	case "C":
		return float64(1000 * hours)
	default:
		return 0.0
	}
}

func main() {
	fmt.Println("Salary calculating mins: 60000, category: 'A'", salary_calculating(60000, "A"))
	fmt.Println("Salary calculating mins: 60000, category: 'B'", salary_calculating(60000, "B"))
	fmt.Println("Salary calculating mins: 60000, category: 'C'", salary_calculating(60000, "C"))
}
