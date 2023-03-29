package main

import (
	"errors"
	"fmt"
)

var err = errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")

func MonthSalary(hours int, value int) (float64, error) {
	if hours < 80 {
		return 0.0, fmt.Errorf("Error: %w", err)
	} else {
		salary := hours * value
		if salary >= 150000 {
			return float64(salary) - (float64(salary) * .10), nil
		} else {
			return float64(salary), nil
		}
	}
}

func main() {
	var salary float64
	var err1 error
	salary, err1 = MonthSalary(90, 150)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Salary: ", salary)
	}
}
