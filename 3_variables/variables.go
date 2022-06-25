package main

import (
	"fmt"
)

// Variables can be declared on package level.
var packageLevelVar1 int

// Multiple variables can be declared at the same time.
var packageLevelVar2, packageLevelVar3, packageLevelVar4 bool

func main() {
	/*
		Variables can be declared at function level the same way
		as on package level.
	*/
	var funcLevelVar1 int

	/*
		Variables declared without an explicit initial value are given
		their zero value. The zero value is:
			- 0 for numeric types,
			- false for the boolean type,
			- the empty string ("") for strings.
	*/

	fmt.Println("Variables with zero values: ")
	fmt.Println("packageLevelVar1:", packageLevelVar1)
	fmt.Println("packageLevelVar2:", packageLevelVar2)
	fmt.Println("packageLevelVar3:", packageLevelVar3)
	fmt.Println("packageLevelVar4:", packageLevelVar4)
	fmt.Println("funcLevelVar1:", funcLevelVar1)

	// Variables can also be initialized with a value.
	var intWithValue1 int = 2

	// When initializing variables with a value, type can be omitted.
	var intWithValue2 = 10

	fmt.Println("Variables initialized with values: ")
	fmt.Println("intWithValue1:", intWithValue1)
	fmt.Println("intWithValue2:", intWithValue2)

	/*
		Inside a function, the := short assignment statement can be used in place
		of a var declaration with implicit type.
	*/
	shortAssignmentVar1 := 4
	fmt.Println("shortAssignmentVar1:", shortAssignmentVar1)

	// Multiple variable initialization.
	var multipleVar1, multipleVar2, multipleVar3 = 10, "hello", true

	fmt.Printf("multipleVar1: type: %T value: %v\n", multipleVar1, multipleVar1)
	fmt.Printf("multipleVar2: type: %T value: %v\n", multipleVar2, multipleVar2)
	fmt.Printf("multipleVar3: type: %T value: %v\n", multipleVar3, multipleVar3)

	// Variable initialization can also be factored like imports.
	var (
		multipleVar4 bool       = false
		multipleVar5 uint64     = 1<<64 - 1
		multipleVar6 complex128 = 10 * 11i
	)

	fmt.Printf("multipleVar4: type: %T value: %v\n", multipleVar4, multipleVar4)
	fmt.Printf("multipleVar5: type: %T value: %v\n", multipleVar5, multipleVar5)
	fmt.Printf("multipleVar6: type: %T value: %v\n", multipleVar6, multipleVar6)

	// Multiple variables can also be assigned with short assignment:
	multipleVar7, multipleVar8 := true, "abc"

	fmt.Printf("multipleVar7: type: %T value: %v\n", multipleVar7, multipleVar7)
	fmt.Printf("multipleVar8: type: %T value: %v\n", multipleVar8, multipleVar8)

	// Values can be converted between compatible types.
	var someInteger1 int = 42
	var someFloat1 float64 = float64(someInteger1)

	fmt.Println("Converted variables:")
	fmt.Println("someInteger1:", someInteger1)
	fmt.Println("someFloat1:", someFloat1)

	/* Constants can be declared the same way as variables.*/
	const someConstant1 int = 42
	fmt.Println("Constants:")
	fmt.Println(someConstant1)
}
