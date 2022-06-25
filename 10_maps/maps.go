package main

import "fmt"

type Coordinate struct {
	Lat, Long float64
}

/*
	A map maps keys to values. The zero value of a map is nil.
	A nil map has no keys, nor can keys be added. The make function
	returns a map of the given type, initialized and ready for use.
*/

func main() {
	// This map has a key of type String, and value of type Coordinate
	var map1 map[string]Coordinate
	map1 = make(map[string]Coordinate)

	map1["Bell Labs"] = Coordinate{
		40.68433, -74.39967,
	}

	fmt.Println("Value from map1:", map1["Bell Labs"])

	// Short initialization works as expected

	map2 := make(map[string]Coordinate)
	fmt.Println("map2:", map2)

	var map3 = map[string]Coordinate{
		"Bell Labs": Coordinate{40.68433, -74.39967},
		"Google":    Coordinate{37.42202, -122.08408},
	}

	// Value types on assignment can be omitted.

	var map4 = map[string]Coordinate{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println("map3:", map3)
	fmt.Println("map4:", map4)

	// Mutating maps

	map5 := make(map[string]int)

	// Inserting new key-value pair
	map5["Answer"] = 42
	map5["Unix"] = 1971

	fmt.Println("map5:", map5)

	// Changing value
	map5["Answer"] = 48

	fmt.Println("map5:", map5)

	// Deleting a key-value pair
	delete(map5, "Answer")

	// Checking if key is present
	value, isPresent := map5["Answer"]
	fmt.Printf("value: %v, isPresent: %v\n", value, isPresent)

}
