package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

// Weird answer for quotients that are not integers
func divide(x int, y int) int {
	return x / y
}

func modulo(x int, y int) int {
	return x % y
}

func main() {
	var operation string
	var x, y int

	fmt.Print("Enter an operation: ")
	fmt.Scan(&operation)

	fmt.Print("Enter a number: ")
	fmt.Scan(&x)

	fmt.Print("Enter another number: ")
	fmt.Scan(&y)

	switch operation {
	case "+":
		fmt.Println(add(x, y))
	case "-":
		fmt.Println(subtract(x, y))
	case "*":
		fmt.Println(multiply(x, y))
	case "/":
		fmt.Println(divide(x, y))
	case "%":
		fmt.Println(modulo(x, y))

	default:
		fmt.Println("Invalid operation")
	}
}
