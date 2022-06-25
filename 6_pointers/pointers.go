package main

import (
	"fmt"
)

/*
	Go has pointers. A pointer holds the memory address of a value.
	The type *T is a pointer to a T value. Its zero value is nil.
	They behave similar to pointers in C, but there is no pointer
	arithmetic.
*/

func main() {
	var i1 int = 40
	var p1 *int = &i1

	fmt.Printf("p1: type %T, value of p1 %v, dereferenced value: %v\n", p1, p1, *p1)

	// Short assignement can also be used like with any other variable.
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
