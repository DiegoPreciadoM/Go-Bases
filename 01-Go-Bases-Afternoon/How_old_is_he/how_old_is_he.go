package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	count_employee := 0
	for _, age := range employees {
		if age > 21 {
			count_employee++
		}
	}

	fmt.Println(count_employee, " empleados son mayores a 21 anios")

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)
}
