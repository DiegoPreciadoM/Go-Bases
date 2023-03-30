package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(ReadData("customers.txt"))
}

func ReadData(filename string) string {
	fi, err := os.Open("customers.txt")

	defer fmt.Println("La ejecucion finalizo")
	defer fi.Close()

	if err != nil {
		panic("El archivo indicado no fue encontrado o esta daniado")
	}

	text, err := io.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	return string(text)
}
