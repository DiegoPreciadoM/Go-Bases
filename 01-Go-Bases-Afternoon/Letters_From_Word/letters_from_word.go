package main

import "fmt"

func main() {

	var word string = "anitalavalatina"

	fmt.Printf("This word \"%s\" has %d letters\n", word, len(word))

	fmt.Println("The letters are:")
	for _, w := range word {
		fmt.Printf("%c\n", w)
	}

}
