package main

import "fmt"

func fibonacciSequenceLoop(index int) int {
	if index == 0 {
		panic("Invalid index")
	}

	if index == 1 {
		return 0
	}

	if index == 2 {
		return 1
	}

	currentValue := 1
	previousValue := 0
	var intermediateValue int

	for i := 3; i <= index; i++ {
		intermediateValue = currentValue

		currentValue = currentValue + previousValue

		previousValue = intermediateValue
	}

	return currentValue
}

func main() {
	var index int

	fmt.Print("Enter an index: ")
	fmt.Scan(&index)

	fmt.Println("Value: ", fibonacciSequenceLoop(index))
}
