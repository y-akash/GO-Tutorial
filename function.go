package main

import (
	"fmt"
	"math"
)

// func functionName() returnType{    if you are not returning something tha dont specify it.
// 	do some action
// }

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// We can pass function inside function
// In cycleNames method we can only pass a function which take one argument of string.
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"cloud", "barret", "tifa"}, sayGreeting)
	cycleNames([]string{"cloud", "barret", "tifa"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)
}
