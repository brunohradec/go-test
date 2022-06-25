package main

/*
	Other modules can be imported using the import keyword. Modules
	can be imported using more separate import statemens like this:

	import "math"
	import "fmt"

	This code groups the imports into a parenthesized, "factored"
	import statement. It is good style to use the factored import
	statements.
*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("The square root of 2 is %g.\n", math.Sqrt(2))
}
