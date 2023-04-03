package main

import "fmt"

func salary_calculating(salary float64) float64 {

	if salary > 150000 {
		return salary - (salary * 0.27)
	} else if salary > 50000 {
		return salary - (salary * 0.17)
	}
	return salary
}

func main() {
	fmt.Println(salary_calculating(30000))
	fmt.Println(salary_calculating(151000))
	fmt.Println(salary_calculating(52000))
}
