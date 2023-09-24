package main

import (
	"fmt"
)

func main() {
	// what is it?
	fmt.Println("Do u elder than 18? (Y/N): ")
	var input string
	fmt.Scanln(&input)

	if input == "Y" {
		addStrangeAction()
	} else {
		fmt.Println("Go away!")
	}
}

func addStrangeAction() {
	fmt.Print("So cool!")
	fmt.Println(" Do u like jokes? (Y/N): ")

	var input string
	fmt.Scanln(&input)

	if input == "Y" {
		fmt.Println("Did you hear about the medieval lamp? It's a knight light. ")
	} else {
		fmt.Println("I'm sorry, I didn't understand you.")
	}
}
