package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func main() {
	var salary int = 100000

	my_error := myError{msg: "el salario ingresado no alcanza el minimo disponible"}

	if salary < 150000 {
		fmt.Println(my_error.Error())
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}
