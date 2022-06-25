package main

import "fmt"

/*
	Go functions may be closures. A closure is a function value that
	references variables from outside its body. The function may access
	and assign to the referenced variables; in this sense the function
	is "bound" to the variables. For example, the adder function returns
	a closure. Each closure is bound to its own sum variable.
*/

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first, second := 0, 1

	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	pos, neg := adder(), adder()

	fmt.Println("Adder:")
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	fib := fibonacci()

	fmt.Println("Fibonacci numbers:")
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
