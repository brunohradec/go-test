package main

import (
	"fmt"
	"time"
)

/*
	Go programs express error state with error values.
	The error type is a built-in interface similar to fmt.Stringer:

	type error interface {
    	Error() string
	}

	As with fmt.Stringer, the fmt package looks for the error interface when
	printing values. Functions often return an error value, and calling code
	should handle errors by testing whether the error equals nil.

	A nil error denotes success; a non-nil error denotes failure.
*/

type MyError struct {
	When time.Time
	What string
}

/*
	Implementation of the string method specified by error interface.
	When MyError happens, fmt will print it like it is specified inside
	this method.
*/
func (e *MyError) Error() string {
	return fmt.Sprintf("at: %v happened: %s", e.When, e.What)
}

/*
	This function will return true if a positive number is even.
	If a negative number is given, it throws value of built-in error
	interface. As MyError struct has a string method from built-in
	error interface implemented, fmt will print error as specified
	by custom method above.
*/
func isPositiveNumberEven(n int) (bool, error) {
	if n >= 0 {
		return n%2 == 0, nil
	} else {
		return false, &MyError{time.Now(), "Given number is negative"}
	}
}

func main() {
	isEven, err := isPositiveNumberEven(2)
	fmt.Printf("Number 2 is even: %v, error: %v\n", isEven, err)

	isEven, err = isPositiveNumberEven(3)
	fmt.Printf("Number 3 is even: %v, error: %v\n", isEven, err)

	isEven, _ = isPositiveNumberEven(2)
	fmt.Printf("Number 2 is even: %v\n", isEven)

	isEven, _ = isPositiveNumberEven(3)
	fmt.Printf("Number 3 is even: %v\n", isEven)

	isEven, err = isPositiveNumberEven(-1)
	fmt.Printf("Number -1 is even: %v, error: %v\n", isEven, err)
}
