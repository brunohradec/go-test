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

/*
	Go's return values may be named. If so, they are treated as variables
	defined at the top of the function. These names should be used to
	document the meaning of the return values. A return statement without
	arguments returns the named return values. This is known as a "naked" return.
	Naked return statements should be used only in short functions, as with
	the example shown here. They can harm readability in longer functions.
*/
func decrementTwo(x, y int) (i, j int) {
	i = x - 1
	j = y - 1
	return
}

func main() {
	fmt.Println("add:", add(1, 2))
	fmt.Println("sub:", sub(4, 2))
	fmt.Println("differentAdd:", differentAdd(4, 2))

	var var1, var2 = 1, 2
	fmt.Println("var1:", var1, "var2:", var2)

	var1, var2 = swap(var1, var2)
	fmt.Println("var1:", var1, "var1:", var2)

	var var3, var4 int = swap(10, 11)
	fmt.Println("var3:", var3, "var4:", var4)

	var5, var6 := swap(20, 40)
	fmt.Println("var5:", var5, "var6:", var6)

	var7, var8 := decrementTwo(10, 11)
	fmt.Println("var7:", var7, "var8:", var8)
}
