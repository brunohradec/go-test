package main

import "fmt"

func add(x int, y int) int {
	return x + 1
}

func sub(x int, y int) int {
	return x - 1
}

/*
	When two or more consecutive named function parameters share a type,
	you can omit the type from all but the last.
*/
func differentAdd(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("add:", add(1, 2))
	fmt.Println("sub:", sub(4, 2))
	fmt.Println("differentAdd:", differentAdd(4, 2))

	var x, y = 1, 2
	fmt.Println("x:", x, "y:", y)

	x, y = swap(x, y)
	fmt.Println("x:", x, "y:", y)

	var i, j int = swap(10, 11)
	fmt.Println("i:", i, "j:", j)

	k, l := swap(20, 40)
	fmt.Println("k:", k, "l:", l)
}
