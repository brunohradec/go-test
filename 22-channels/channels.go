package main

import "fmt"

/*
	Channels are a typed conduit through which you can send and receive values
	 with the channel operator, <-.

	ch <- v    // Send v to channel ch.
	v := <-ch  // Receive from ch, and assign value to v.

	The data flows in the direction of the arrow.

	Like maps and slices, channels must be created before use:

	ch := make(chan int)

	By default, sends and receives block until the other side is ready.
	This allows goroutines to synchronize without explicit locks or
	condition variables.
*/

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	// Send the current sum to channel c
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Creating an integer channel
	c := make(chan int)

	// Starting the sum function as goroutine for the first half of the slice.
	go sum(s[:len(s)/2], c)

	// Starting the sum function as goroutine for the second half of the slice.
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("x+y:", x+y)

	/*
		Channels can be buffered. Provide the buffer length as the second argument
		to make to initialize a buffered channel. Sends to a buffered channel block
		only when the buffer is full. Receives block when the buffer is empty.
	*/

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println("Buffered channel:")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
