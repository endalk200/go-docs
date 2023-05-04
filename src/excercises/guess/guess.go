package main

import (
	"fmt"
	"math/rand"
)

func guessName() {
	potentialNames := []string{
		"James",
		"John",
		"Britney",
		"Robert",
		"Mary",
		"Michael",
		"William",
		"David",
	}

	var correctGuess = potentialNames[rand.Intn(len(potentialNames))]
	var guessedName string

	fmt.Print("Enter your guess: ")
	fmt.Scan(&guessedName)

	if guessedName == correctGuess {
		fmt.Println("You guessed right")
	} else {
		fmt.Printf("You guessed wrong. Correct value is %s\n", correctGuess)
	}
}

func guessNumber() {
	var correctGuess int = rand.Intn(10)

	var guessedNumber int

	fmt.Print("Enter your guess: ")
	fmt.Scan(&guessedNumber)

	if correctGuess == guessedNumber {
		fmt.Println("You guessed right")
	} else {
		fmt.Printf("You guessed wrong. Correct value is %d\n", correctGuess)
	}
}

func main() {
	var whatToGues string

	fmt.Print("What do you want to guess? number/name: ")
	fmt.Scan(&whatToGues)

	switch whatToGues {
	case "number":
		guessNumber()
	case "name":
		guessName()
	default:
		fmt.Println("Invalid choice")
	}
}
