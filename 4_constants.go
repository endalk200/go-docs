package main

import (
	"fmt"
	"math"
	"reflect"
)

// const declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 500000000

	println("Typeof const n", reflect.TypeOf(n))
	println("Typeof const s", reflect.TypeOf(s))

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(int64(d))

	/** A number can be given a type by using it in a context that requires one,
	such as a variable assignment or function call. For example,
	here math.Sin expects a float64.
	*/
	fmt.Println(math.Sin(n))
}
