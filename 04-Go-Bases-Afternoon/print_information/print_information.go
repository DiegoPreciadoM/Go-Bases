package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	defer fmt.Println("Ejecucion Finalizada")
	fmt.Println(ReadData("customers.txt"))
}

func ReadData(filename string) string {

	file, err := os.Open(filename)

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta daniado")
	}

	text, err2 := io.ReadAll(file)
	if err2 != nil {
		panic("El archivo indicado no fue encontrado o esta daniado")
	}

	defer file.Close()

	return string(text)
}
