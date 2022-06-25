package main

import "fmt"

/* A struct is a collection of fields. */

type Point struct {
	X int
	Y int
}

func main() {
	fmt.Println("New point:", Point{1, 2})

	// Asignment with zero values
	var point1 Point
	fmt.Println("point1:", point1)

	// Assignment with values
	var point2 = Point{1, 2}
	fmt.Println("point2:", point2)

	// Short assignment
	point3 := Point{1, 2}
	fmt.Println("point3:", point3)

	// Changing values
	point3.X = 10
	fmt.Println("point3:", point3)

	/*
		To access the field X of a struct when we have the struct pointer p1
		we could write (*p1).X. However, that notation is cumbersome, so the
		language permits us instead to write just p.X, without the explicit
		dereference.
	*/

	point4 := Point{1, 2}
	p1 := &point4
	p1.X = 100
	fmt.Println("point4:", point4)

	// Different ways to assign structures.

	point5 := Point{1, 2} // has type Point
	point6 := Point{X: 1} // Y:0 is implicit
	point7 := Point{}     // X:0 and Y:0
	p2 := &Point{1, 2}    // has type *Point

	fmt.Println("point5:", point5)
	fmt.Println("point6:", point6)
	fmt.Println("point7:", point7)
	fmt.Printf("p2: type: %T, value: %p, dereferenced value: %v\n", p2, p2, *p2)
}
