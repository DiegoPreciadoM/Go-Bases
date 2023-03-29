package main

import (
	"errors"
	"fmt"
)

type myError struct {
	err2 error
}

var err1 = errors.New("el salario es menor a 10000")

func main() {
	var salary int = 10000

	my_error := myError{err2: fmt.Errorf("Error: %w", err1)}
	coincidence := false
	var err error

	if salary <= 10000 {
		err = my_error.err2
	}

	coincidence = errors.Is(err1, errors.Unwrap(err))

	if coincidence {
		fmt.Println(err)
	} else {
		fmt.Println("True")
	}
}
