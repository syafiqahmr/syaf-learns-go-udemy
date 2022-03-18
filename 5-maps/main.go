package main

import "fmt"

// map <-> dict in python
// maps is a series of key value pair
// all they keys must be of the same type
// all the values must be of the same type

func main() {
	// ways to declare a map -
	// syntax 1:
	// note how theres a comma at the end
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4b745",
		"white": "#fffff",
	}
	// prints map[green:#4b745 red:#ff000 white:#fffff]
	fmt.Println(colors)

	// syntax2 :
	var colors2 map[string]string
	// prints map[] aka the 0 value
	fmt.Println(colors2)

	// syntax 3:
	// similar method to 2
	colors3 := make(map[string]string)
	fmt.Println(colors3)

	// add values to map
	colors3["white"] = "fffff"
	fmt.Println(colors3)

	// delete values
	delete(colors3, "white")
	fmt.Println(colors3)

	printMap(colors)
}

// recall that map is reference type
// so we can simply do this and not care abt pointers
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// MAP VS STRUCT

// map
// 1. all keys and must be the same type
// 2. keys are indexed - we can iterate over them
// 3. reference type
// 4. used to represent a collection of closely related property
// 5. you do not need to know all the keys at compile time

// struct
// 1. values can be of different type
// 2. keys does not support indexing
// 3. value type
// 4. used to represent a "thing" with a lot of different properties
// 5. you need to know all the different fields at compile time
