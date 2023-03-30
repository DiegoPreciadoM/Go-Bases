package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando...")
	_, err := os.Open("customers.txt")
	if err != nil {
		defer func() {
			fmt.Println("Ejecucion Finalizada")
		}()
		panic("El archivo indicado no fue encontrado o esta daniado")
	}
	fmt.Println("Ejecucion Finalizada")
}
