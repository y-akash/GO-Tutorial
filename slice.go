package main

import "fmt"

func main() {
	// slices (use arrays under the hood)
	// here we are not giving the size so it becomes slice
	var scores = []int{100, 50, 60}
	scores[2] = 25
	// we can append value in slice using append(sliceVariable, value) method
	// but this method returns a new slice it dose't append in the passed slice
	// thats why we are storing in the existing slice so that it will update
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	// slice ranges
	// in this we get a range of element from slice or array and stored it as a new slice
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
