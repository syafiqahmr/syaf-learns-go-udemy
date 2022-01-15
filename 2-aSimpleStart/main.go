// package = project = workspace
// package is a packet of common source code files
// go has 2 types of pacakages:
// 	1. executable -> generates a file the we can run
//	2. reusable -> code used as helpers (eg. libraries)

// specifically, if you name your pacakge 'main',
// running go build will create an executable file
// other package names will be treated as reusable package

// all files in a main package needs to have a main func

// all the files in a package needs to have this as the first line
package main

// fmt refers to format
// fmt is a standard library included in Go by default
// you can also import other reusable packages
// golang.org/pkg for the official docs
import "fmt"

// func = function
func main() {
	fmt.Println("Hi there!")
}

// main.go file is typically organised like this:
// 1. package name
// 2. import stmts
// 3. func main(){} -> this is automatically run when executed
