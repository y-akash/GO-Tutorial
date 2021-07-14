package main

import "fmt"

func main() {

	// var variableName [sizeOfArray]TypeOfvalues = [sizeOfArray]TypeOfvalues{val1, val2, val3}
	var ages [3]int = [3]int{20, 25, 30}
	var ages1 = [3]int{20, 25, 30}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	names[1] = "luigi"

	fmt.Println(ages)
	fmt.Println(ages1, len(ages1))
	fmt.Println(names, len(names))

}
