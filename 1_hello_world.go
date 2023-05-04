package main

import "fmt"

func main() {
	fmt.Println("hello world")                                           // Include a new line and include a space
	fmt.Print("hello world")                                             // No new line and no space
	fmt.Printf("\nhello world! My name is %s. I love go: %d", "John", 1) // No new line and no space

	// fmt.Fprint() - write to an external location
}
