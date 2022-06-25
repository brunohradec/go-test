package main

import "fmt"

/*
	One of the most ubiquitous interfaces is Stringer defined by the fmt package.

	type Stringer interface {
		String() string
	}

	A Stringer is a type that can describe itself as a string. The fmt package
	(and many others) look for this interface to print values. It is used to
	make custom types printable by the fmt package.
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Dennis Ritchie", 71}
	z := Person{"Ken Thompson", 71}
	fmt.Println(a, z)
}
