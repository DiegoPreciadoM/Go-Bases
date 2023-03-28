package main

import "fmt"

func main() {
	var age int = 22
	var is_working bool = true
	var antiquity int = 1
	var salary float64 = 300000.00

	if age < 22 {
		fmt.Println("We cant give you a Lending")
	} else if age >= 22 && is_working && antiquity >= 1 {
		if salary > 100000 {
			fmt.Println("Lending for you, with out taxes")
		} else {
			fmt.Println("Lending for you, with taxes")
		}
	}
}
