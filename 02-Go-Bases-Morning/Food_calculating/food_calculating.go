package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Dog(dog_1 int) int {
	return 10000 * dog_1
}

func Cat(cat_1 int) int {
	return 5000 * cat_1
}

func Hamster(hamster_1 int) int {
	return 250 * hamster_1
}

func Tarantula(tarantula_1 int) int {
	return 150 * tarantula_1
}

func Animal(animal string) func(int) int {
	switch animal {
	case dog:
		return Dog
	case cat:
		return Cat
	case hamster:
		return Hamster
	case tarantula:
		return Tarantula
	default:
		return nil
	}
}

func main() {
	dog_1 := Animal(dog)
	cat_1 := Animal(cat)
	hamster_1 := Animal(hamster)
	tarantula_1 := Animal(tarantula)

	fmt.Println("Dog: ", dog_1(2))
	fmt.Println("Cat: ", cat_1(5))
	fmt.Println("Hamster: ", hamster_1(2))
	fmt.Println("Tarantula: ", tarantula_1(5))
}
