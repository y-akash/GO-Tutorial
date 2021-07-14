package main

import "fmt"

func main() {
	// string variables
	var nameOne string = "mario" // var VariableName type(optional)
	var nameTwo = "luigi"        // The above is also same with this line but here we dont have to explicitly tell the type
	var nameThree string         //default value will be empty string

	fmt.Print(nameOne, nameTwo, nameThree)

	nameOne = "peach" // we cant change to another type of variable once it is declared.
	nameThree = "bowser"

	fmt.Print(nameOne, nameTwo, nameThree)

	// This is a short way to write the variable but it is not allowed outside of the method.
	// it similar to var nameOne string = "mario"
	nameFour := "yoshi"
	fmt.Print(nameFour)
}
