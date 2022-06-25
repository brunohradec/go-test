package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	X, Y float64
}

/*
	Go does not have classes. However, you can define methods on types.
	A method is a function with a special receiver argument.
	The receiver appears in its own argument list between the func
	keyword and the method name.
*/

func (c Coordinate) Abs() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

/*
	You can declare a method on non-struct types, too.
	In this example we see a numeric type MyFloat with an Abs method.
	You can only declare a method with a receiver whose type is defined
	in the same package as the method. You cannot declare a method with
	a receiver whose type is defined in another package, which includes
	the built-in types such as int.
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	c := Coordinate{3, 4}
	fmt.Println(c.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
