package main

import "fmt"

/*
	The type [n]T is an array of n values of type T.
	The expression var a [10]int declares a variable a as an array of ten integers.
	An array's length is part of its type, so arrays cannot be resized. This seems
	limiting, but don't worry; Go provides a convenient way of working with arrays.
*/

func main() {
	var array1 [2]int

	// Without initialization, every element is set to it's zero values
	fmt.Println("Zero values:")
	fmt.Println("array1[0]", array1[0])
	fmt.Println("array1[1]", array1[1])
	fmt.Println("array1:", array1)

	// Changing elements is same as in C.
	array1[0] = 2
	array1[1] = 4

	fmt.Println("After initialization")
	fmt.Println("array1[0]", array1[0])
	fmt.Println("array1[1]", array1[1])
	fmt.Println("array1:", array1)

	// Arrays can be initialized on declaration.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:", primes)

	// Uninitialized memebers are set to their zero values.
	array2 := [10]int{1, 2}
	fmt.Println("array2:", array2)
}
