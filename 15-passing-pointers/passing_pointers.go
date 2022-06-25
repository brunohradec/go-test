package main

import "fmt"

/*
	FUNCTIONS AND METHODS DO NOT DO PERMANENT CHANGES ON VARIABLES PASSED
	BY VALUE, THEY RETURN A CHANGED COPY INSTEAD. WHEN AN ARGUMENT IS PASSED
	BY VALUE, A COPY IS MADE ON FUNCTION STACK, AND EVERYTHING THAT HAPPENS INSIDE
	A FUNCTION IS DONE ON THAT COPY.

	WHEN AN ADDRESS IS PASSED AS A POINTER, CHANGES CAN BE MADE ON THE ORIGINAL
	VALUE.
*/

// Function

func Increment(i *int) {
	*i += 1
}

// Method

type Coordinate struct {
	X, Y float64
}

func (c *Coordinate) Scale(f float64) {
	/*
		REMEMBER THAT POINTERS TO STRUCTS DO NOT NEED TO BE DEREFERENCED AS
		GO DOES AUTOMATIC DEREFERENCING WITH STRUCTS (SYNTAX SUGAR)!
	*/

	c.X = c.X * f
	c.Y = c.Y * f
}

func main() {
	int1 := 10

	coordinate1 := Coordinate{1, 2}
	coordinate2 := Coordinate{3, 4}

	Increment(&int1)

	/*
		When using methods, both calling on pointer or original struct is the same.
		The same way Go does automatic dereferencing on member acces for pointers to
		structs, it does automatic conversion to pointer on method calls.
	*/
	coordinate1.Scale(10)
	p := &coordinate2
	p.Scale(20)

	fmt.Println("int1", int1)
	fmt.Println("coordinate1", coordinate1)
	fmt.Println("coordinate2", coordinate2)
}
