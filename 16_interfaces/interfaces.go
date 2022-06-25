package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var abser1 Abser

	float1 := MyFloat(-math.Sqrt2)
	vertex1 := Vertex{3, 4}

	abser1 = float1 // a MyFloat implements Abser

	/*
		Calling a method on an interface value executes the method
		of the same name on its underlying type.
	*/
	fmt.Println(abser1.Abs()) // MyFloat Abs is called

	abser1 = &vertex1 // a *Vertex implements Abser

	fmt.Println(abser1.Abs()) // *Vertex Abs is called

	// In the following line, vertex1 is a Vertex (not *Vertex) and does NOT implement Abser:
	// abser1 = vertex1 THIS IS AN ERROR!

	/*
		A type assertion provides access to an interface value's underlying concrete value.
		Statement t := i.(T) asserts that the interface value i holds the concrete type T
		and assigns the underlying T value to the variable t.

		If i does not hold a T, the statement will trigger a panic.

		To test whether an interface value holds a specific type, a type assertion can return
		two values: the underlying value and a boolean value that reports whether the assertion
		succeeded.

		Example: t, ok := i.(T)

		If i holds a T, then t will be the underlying value and ok will be true.
		If not, ok will be false and t will be the zero value of type T, and no
		panic occurs.
	*/

	abser1 = float1
	float2 := abser1.(MyFloat)

	fmt.Printf("Underlying MyFloat value: %v, type: %T\n", float2, float2)

	float2, isMyFloat := abser1.(MyFloat)
	fmt.Printf("Underlying MyFloat value: %v, isMyFloat: %v\n", float2, isMyFloat)
}
