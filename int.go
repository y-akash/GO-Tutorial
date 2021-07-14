package main

import "fmt"

func main() {
	// int variables
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Print(ageOne, ageTwo, ageThree)

	// bits & memory
	// intBit means that the range of value
	// int8 -128 to 127
	var numOne int8 = 25
	var numTwo int8 = 124  // too large a number for 8-bit
	var numThree uint = 25 // uint means we only use non-negative numbers
	var numFour uint8 = 25 // uint means we only use non-negative numbers with bit limitaions

	fmt.Print(numOne, numTwo, numThree, numFour)

	// but for float we have to give bit, its a cumplsory but not in int
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64

	fmt.Print(scoreOne, scoreTwo, scoreThree)

	// for more info see https://golang.org/ref/spec#Numeric_types
}
