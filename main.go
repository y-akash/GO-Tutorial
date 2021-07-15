package main

import "fmt"

// the scope of this variable is package scope
var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func main() {
	// Here we are able to access sayHello() method of greeting.go file
	// bcoz the package name of both file are same.
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}
