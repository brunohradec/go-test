package main

import "fmt"

/*
	An empty interface can be used when there is code independent of types,
	as every type matches the empty interface.s
*/

func do(i interface{}) {
	/*
		A type switch is like a regular switch statement, but the cases in a
		type switch specify types (not values), and those values are compared
		against the type of the value held by the given interface value.
	*/

	switch v := i.(type) {
	case int:
		// Type specific operations can be done inside the matching case.
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
