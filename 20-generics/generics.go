package main

import "fmt"

/*
	Go functions can be written to work on multiple types using type parameters.
	The type parameters of a function appear between brackets, before the
	function's arguments.
*/

/*
	Interfaces are used as type constraints. Interface any is an empty interface
	accepting all types, but other interfaces like built-in comparable interface
	or custom interfaces can be used as type constraints.
*/

func printOne[T any](x T) {
	fmt.Println(x)
}

// A custom interface that can be used as type constraint

type number interface {
	int | float64
}

func add[T number](a T, b T) T {
	return a + b
}

func incrementTwo[A number, B number](a *A, b *B) {
	*a += 1
	*b += 1
}

func main() {
	var int1 int = 1
	var float1 float64 = 3.14

	fmt.Println(add(int1, int1))
	incrementTwo(&int1, &float1)

	printOne(int1)
	printOne(float1)
}
