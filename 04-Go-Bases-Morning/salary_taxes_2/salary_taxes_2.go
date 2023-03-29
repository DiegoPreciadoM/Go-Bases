package main

import (
	"errors"
	"fmt"
)

type myError struct{}

var err1 = errors.New("el salario es menor a 10000")

func (e *myError) Error() error {
	return fmt.Errorf("Error: %w", err1)
}

func main() {
	var salary int = 10000

	my_error := myError{}
	var err error
	coincidence := false

	if salary <= 10000 {
		err = my_error.Error()
		//err = errors.New("Error: el salario es menor a 10000")
	}

	coincidence = errors.Is(err, err1)

	if coincidence {
		fmt.Println(err)
	} else {
		fmt.Println("TRUE")
	}
}
