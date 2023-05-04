package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("'Hello' has type of: ", reflect.TypeOf("Hello"))
	fmt.Println("'2' has type of: ", reflect.TypeOf(2))
	fmt.Println("'true' has type of: ", reflect.TypeOf(true))
}
