package main

import (
	"fmt"
)

// struct is something like dict in python
// 1. define the fields the struct has
type contactInfo struct {
	// property name, type
	// note there is no comma
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// you can embed another struct into a struct
	contact contactInfo
}

type person2 struct {
	firstName string
	lastName  string
	// you can choose to not type a name for contact info
	contactInfo
}

func main() {
	contact := contactInfo{email: "alex@gmail.com", zipCode: 101000}

	// 2. declare the struct
	// 1st syntax: note how you do not need to specify
	// that Alex is the first name and Anderson is the last name
	alex := person{"Alex", "Anderson", contact}
	// 2nd syntax
	alex2 := person{firstName: "Alex", lastName: "Anderson", contact: contact}
	// 3rd syntax
	// go will automatically assign a zero value for the fields
	// zero values: string - "", int/float - 0, bool - false
	var alex3 person
	// 4th syntax
	// note how there is a comma at the end
	alex4 := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 101000,
		},
	}

	// note the field name of contactInfo
	alex5 := person2{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 101000,
		},
	}

	// updating fields
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// prints {Alex Anderson {alex@gmail.com 101000}}
	fmt.Println(alex)
	fmt.Println(alex2)
	fmt.Println(alex4)
	fmt.Println(alex5)
	// prints { { 0}}
	fmt.Println(alex3)
	// prints {firstName: lastName: contact:{email: zipcode:0}}
	fmt.Printf("%+v", alex3)
	alex3.print()

	alex.updateNameWrong("Alexy")
	// note it still prints Alex as first name
	alex.print()

	// & is an operator
	// &variable : give me the memory address of the value this var is pointing at
	// note it does NOT point to the struct itself
	// but instead the address in memory
	alexPointer := &alex
	alexPointer.updateName("Alexy")
	alex.print()

	// shortcut done by go which does the above:
	// even though the type here is person
	// while updateName is expecting *person
	// go allows updateName to accept *person or person
	alex.updateName("Alex")
}

// struct can be a receiver too
func (p person) print() {
	fmt.Printf("%+v", p)
}

// note that go is pass by value language
// whenever u pass a value into a function
// the value will be placed into a new container in the memory
// so the update is not done on the original struct
// but instead on the copy
func (p person) updateNameWrong(newFirstName string) {
	p.firstName = newFirstName
}

// * here is a type description
// it means we're working with a pointer to a person
// note that everything in Go is passed by value
// so if you print &pointerToPerson it will point to a diff address
func (pointerToPerson *person) updateName(newFirstName string) {
	// * here is an operator
	// give me the value this memory address is pointing at
	// so it refers to the struct
	(*pointerToPerson).firstName = newFirstName

	// summary
	// in memory:
	// address		value
	// 0001				person(firstName:"Alex"..)
	// turn address into value with *address
	// turn value into address wirth &value

	// if we see a * before a type
	// meanwhile the * as an operator is in front of a variable/pointer
	// that means we're looking for a type
	// that is a pointer to a person

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	// slice here is actl updated
	fmt.Println(mySlice)

	// array vs slice
	// primitive data structure vs can grow and shrink
	// can't be resized vs can be resized
	// rarely used directly vs used 99% the time for lists of elements

	// sllice is a reference type
	// which means it refers to another ds in memory
	// when you first initialise a slice, you create a slice ds
	// and an array ds with the things you initialise
	// the first element in the slice will point to the array
	// then as you add more items to the slice
	// the slice looks like this:
	// a structure that records the length of slice, the capacity of the slice,
	// and a reference to the underlying array
	// while the array looks like:
	// Hi, There, How ...

	// so when you pass the slice through a function
	// you are copying the slice
	// but the pointer remains the same
	// and you still point to the same array created earlier

	// value types:
	// int, float, string, bool, structs
	// use pointers to change these things in a function

	// reference types:
	// slices, maps, channels, pointers, functions
	// don't worry about pointers with these
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
