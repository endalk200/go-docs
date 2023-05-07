# Data types in go

## Arrays

Arrays can be defined in Golang as below.

```go
arr1 := [3]int {7,9,4}
// array size is 3
```

This is the easiest way to declare an Array and assign its values in Golang. **But this method can only be used when you initialize a variable inside a function.** Other than this, you can follow two other ways to declare an Array in your Go code.

```go
// ex 1:
var arr1 [3]int = [3]int {2,3}

// ex 2:
var arr1 = [3]int {1,7,9}
```

When you declare an Array, **you need to clearly define its size and when you initialize the Array, you can assign values up to that size.**
Once you define and set up Array values, you can easily print them using the “fmt”.

```go
// ex:
fmt.Println(arr1)
```

When you need to change Array values, you can use the relevant index and change its value accordingly.

```go
// ex:
arr1[2] = 6
fmt.Println(arr1)
```

For more array examples, please refer to the this [array examples in go](../8_arrays.go).

## Slices

You can use the same method which we used to declare an Array to declare a Slice in Golang. But the difference is, **in slices you don’t have to mention the size for that particular slice.** But in Arrays, we had to define its size when declaring.

```go
cities := []string {"London", "NYC", "Colombo", "Tokyo"}

// OR

var cities = []string {"London", "NYC", "Colombo", "Tokyo"}

// OR

var cities []string = []string {"London", "NYC", "Colombo", "Tokyo"}
```

You can use any method from the above list when you declare a Slice in your Go code. Also, you can easily print a slice using “fmt” as below.

```go
fmt.Println(cities)
```

Since a slice doesn’t have a specific size, we can append values to a slice in addition to changing its existing values. But in Arrays, we couldn’t append values and all we could do is change its existing values.

You can append values to a slice using two different methods. If you want to append values to the original slice, you can do it as below.

```go
cities = append(cities, "LA")
fmt.Println(cities)
```

This will add LA to the end of the original slice and update it. Then, when you print the original slice, it will be outputted LA along with other city names.
But in slices, you can use the built-in append() function to add values and it will be creating another slice and append values without overwriting the original slice.

```go
// ex:
cities := []string {"London", "NYC", "Colombo", "Tokyo"}
fmt.Println(cities)

addCity := append(cities, "Auckland")
fmt.Println(addCity)
fmt.Println(cities)
```

In the above example, I created a separate variable called addCity and used the append function to add values. When you print the addCity variable, it will output cities along with Auckland but it will not overwrite the original slice. You can verify it by printing the original slice once again.

```
// output
[London NYC Colombo Tokyo]
[London NYC Colombo Tokyo Auckland]
[London NYC Colombo Tokyo]
```

Also, you can use this append function directly with the fmt.Println() and it will also do the same as above.

```go
fmt.Println(append(cities, "Auckland"))
```

For more array examples, please refer to the this [slice examples in go](../9_slices.go).
