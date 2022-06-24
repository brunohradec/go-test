package main

import "fmt"

/*
	An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
	flexible view into the elements of an array. In practice, slices are much more
	common than arrays.

	The type []T is a slice with elements of type T. A slice is formed by specifying
	two indices, a low and high bound, separated by a colon: a[low : high]

	This selects a half-open range which includes the first element, but excludes the
	last one. The following expression creates a slice which includes elements 1 through
	3 of a: a[1:4]
*/

func main() {
	array1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := array1[0:4]
	fmt.Println("slice1:", slice1)

	/*
		A slice does not store any data, it just describes a section of an
		underlying array. Changing the elements of a slice modifies the
		corresponding elements of its underlying array. Other slices that
		share the same underlying array will see those changes.
	*/

	slice2 := array1[2:4]

	slice1[0] = 20 // slice1[0] is array1[0]
	slice2[0] = 40 // slice2[0] is array1[2]

	fmt.Println("After change:")
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("array1:", array1)

	/*
		A slice has both a length and a capacity.
		The length of a slice is the number of elements it contains.
		The capacity of a slice is the number of elements in the underlying
		array, counting from the first element in the slice.
	*/

	fmt.Println("slice1 length:", len(slice1))
	fmt.Println("slice1 capacity", cap(slice1))

	/*
		When slicing, you may omit the high or low bounds to use their defaults instead.
		The default is zero for the low bound and the length of the slice for the high bound.
	*/

	slice3 := array1[:8]
	slice4 := array1[2:]
	slice5 := array1[:]

	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)
	fmt.Println("slice5:", slice5)

	// Slices can be resliced if there is enough capacity.

	// Slice the slice to give it zero length.
	slice1 = slice1[:0]
	fmt.Println("slice1[:0]:", slice1)

	// Extend its length.
	slice1 = slice1[:8]
	fmt.Println("slice1[:8]:", slice1)

	// Drop its first two values. This does not resize the slice as it is in the bounds!
	slice1 = slice1[2:]
	fmt.Println("slice1[2:]:", slice1)

	/*
		A slice literal is like an array literal without the length.
		This is an array literal: [3]bool{true, true, false}
		And this creates the same array as above, then builds a slice
		that references it: []bool{true, true, false}
	*/

	slice6 := []int{1, 2, 3, 4}
	slice7 := []bool{true, false, false, true}

	fmt.Println(slice6)
	fmt.Println(slice7)

	// The zero value of a empty silce is nil.
	var slice8 []int

	fmt.Println("slice8:", slice8)
	fmt.Println("slice8 == nil:", slice8 == nil)
}
