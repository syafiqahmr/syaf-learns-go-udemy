package main

import "fmt"

// to all the custom types
type bot interface {
	// if you are a type in this program with
	// a function called getGReeting,
	// you are now a member of type bot
	// meed to list argument types and return types
	// eg. getGreeting (string, int) (string, error)
	// in the case below, no arguments but expects to return string
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// since we are not using englishBot
// can change from func (eb englishBot)
// to just func (englishBot)
func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

// all members of bot can now call this function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// concrete vs interface type:
// concrete types - map, struct, englishBot, ...
// interface types - bot
// concrete types mean you can assign a value to it

// other things to note regarding interfaces
// 1. interfaces are not generic types (like java)
// 2. interfaces are implicit - you do not need to explicitly link englishBot and Bot
// 3. interfaces are a contract to help us manage types
// 4. interfaces are tough and first, we need to understand
// how to read interfaces in the standard lib
