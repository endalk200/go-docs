package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	// var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	/**Variables declared without a corresponding initialization
	are zero-valued. For example, the zero value for an int is 0.*/
	var e int // 0
	// var e float32 // 0
	// var e string // ""
	// var e bool // false
	fmt.Println(e)
	fmt.Println("TypeOf e: ", reflect.TypeOf(e))

	/**The := syntax is shorthand for declaring and initializing
	a variable, e.g. for var f string = "apple" in this case.
	This syntax is only available inside functions.
	In Go, := is for declaration + assignment, whereas = is for assignment only.
	*/
	f := "apple" // var f string = "apple"
	fmt.Println(f)
}
