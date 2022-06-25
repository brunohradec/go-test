package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sum := 0

	/*
		Go has only one looping construct, the for loop.
		The basic for loop has three components separated by semicolons:

			- the init statement: executed before the first iteration
			- the condition expression: evaluated before every iteration
			- the post statement: executed at the end of every iteration

		The init statement will often be a short variable declaration, and
		the variables declared there are visible only in the scope of the for statement.

		The loop will stop iterating once the boolean condition evaluates to false.
	*/

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// The init and post statements are optional.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	/*
		Go's if statements are like its for loops; the expression need
		not be surrounded by parentheses but the braces are required
	*/

	var1, var2 := 10, 20

	if var1 < var2 {
		fmt.Println("var1 is smaller than var2.")
	}

	/*
		Like for, the if statement can start with a short statement to
		execute before the condition. Variables declared by the statement
		are only in scope until the end of the if.
	*/

	if var3 := 100; var1 < var3 {
		fmt.Println("var1 is smaller than var3.")
	}

	/*
		Variables declared inside an if short statement are also available
		inside any of the else blocks.
	*/

	if var4 := 4; var1 < var4 {
		fmt.Println("var1 is smaller than var3.")
	} else {
		fmt.Println("Value of var4 is", var4)
	}

	/*
		A switch statement is a shorter way to write a sequence of if - else
		statements. It runs the first case whose value is equal to the condition expression.
	*/

	os := "mac OS"

	switch os {
	case "mac OS":
		fmt.Println("Mac!")
	case "Linux":
		fmt.Println("Linux!")
	default:
		fmt.Println("I don't know this os.")
	}

	// Like all other statements, swith also supports the init part.

	switch realOS := runtime.GOOS; realOS {
	case "darwin":
		fmt.Println("mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	// Switch does7nt have to have a condition.

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	/*
		A defer statement defers the execution of a function until
		the surrounding function returns.
	*/
	defer fmt.Println("The main function has finished executing.")

	/*
		Deferred function calls are pushed onto a stack. When a function returns,
		its deferred calls are executed in last-in-first-out order.
	*/
	defer fmt.Println("I was after first defer.")
}
