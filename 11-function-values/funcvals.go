package main

import "fmt"

/*
	Functions are values too. They can be passed around just like other values.
	Function values may be used as function arguments and return values.
*/

func doMath(val1, val2 int, fn func(int, int) int) int {
	return fn(val1, val2)
}

func add(val1, val2 int) int {
	return val1 + val2
}

func sub(val1, val2 int) int {
	return val1 - val2
}

func main() {
	fmt.Println("Addition:", doMath(1, 2, add))
	fmt.Println("Subtraction:", doMath(1, 2, sub))
}
